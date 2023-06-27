package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const url = "https://groceries.asda.com/sitemap"
const driverBin = "./bin/chromedriver"
const port = 4444

type Link struct {
	URL  string
	Text string
}

func main() {
	service, err := selenium.NewChromeDriverService(driverBin, port)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--headless",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	// Go to site map
	log.Printf("Going to %s\n", url)
	driver.Get(url)

	// Wait until the last item, identified by XPATH, is loaded.
	var lastItem []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(50 * time.Millisecond)
		lastItem, _ = driver.FindElements(selenium.ByXPATH, "/html/body/div[1]/div[2]/section/main/div[4]/div/div[50]/div[4]/div[2]/li/a/span")
		return len(lastItem) > 0, nil
	})

	// Get all anchor tags in the site map
	anchors, err := driver.FindElements(selenium.ByTagName, "a")
	if err != nil {
		log.Fatal(err)
	}

	shelvesAisles := []Link{}

	// Iterate over `a` tags and collect the aisle and shelf links.
	for _, a := range anchors {
		href, err := a.GetAttribute("href")
		if err != nil {
			log.Println(err)
			continue
		}

		if strings.Contains(href, "/shelf/") || strings.Contains(href, "/aisle/") {
			txt, _ := a.Text()
			log.Printf("Adding shelf/aisle Link %s\n", href)
			shelvesAisles = append(shelvesAisles, Link{
				URL:  href,
				Text: txt,
			})
		}
	}

	productPages := []Link{}

	for _, shelfAisle := range shelvesAisles {
		if strings.Contains(shelfAisle.Text, "View All") {
			continue
		}

		log.Printf("Going to %s\n", shelfAisle.URL)
		driver.Get(shelfAisle.URL)

		// Wait for main product list to load.
		var mainProductDiv selenium.WebElement
		driver.Wait(func(wd selenium.WebDriver) (bool, error) {
			time.Sleep(50 * time.Millisecond)
			mainProductDiv, _ = driver.FindElement(selenium.ByClassName, "co-product-list__main-cntr")
			return mainProductDiv != nil, nil
		})

		// Get all anchor tags inside main product list.
		var productAnchors []selenium.WebElement
		driver.Wait(func(wd selenium.WebDriver) (bool, error) {
			time.Sleep(50 * time.Millisecond)
			productAnchors, _ = mainProductDiv.FindElements(selenium.ByTagName, "a")
			return len(productAnchors) > 0, nil
		})

		// Iterate over anchors and get the urls.
		for _, a := range productAnchors {
			href, err := a.GetAttribute("href")
			if err != nil {
				log.Println(err)
				continue
			}

			txt, _ := a.Text()
			log.Printf("Adding product page %s\n", href)
			productPages = append(productPages, Link{
				URL:  href,
				Text: txt,
			})
		}
	}

	fmt.Println(len(productPages))

	file, _ := json.MarshalIndent(productPages, "", " ")
	_ = os.WriteFile("./productlinks.json", file, 0644)
}
