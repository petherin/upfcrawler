package asda

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
const driverBin = "bin/chromedriver"
const port = 4444
const shelfLink = "/shelf/"
const aisleLink = "/aisle/"
const aTag = "a"
const hrefTag = "href"
const productLinksFile = "./productlinks.json"
const mainProductClass = "co-product-list__main-cntr"
const lastElementXPath = "/html/body/div[1]/div[2]/section/main/div[4]/div/div[50]/div[4]/div[2]/li/a/span"

type Link struct {
	URL  string
	Text string
}

func Crawl() {
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
		lastItem, _ = driver.FindElements(selenium.ByXPATH, lastElementXPath)
		return len(lastItem) > 0, nil
	})

	// Get all anchor tags in the site map
	anchors, err := driver.FindElements(selenium.ByTagName, aTag)
	if err != nil {
		log.Fatal(err)
	}

	shelvesAisles := []Link{}

	// Iterate over `a` tags and collect the aisle and shelf links.
	for _, a := range anchors {
		href, err := a.GetAttribute(hrefTag)
		if err != nil {
			log.Println(err)
			continue
		}

		if strings.Contains(href, shelfLink) || strings.Contains(href, aisleLink) {
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
		// Links are organised by:
		// Cat (Summer)
		// 		Dept (BBQ)
		//			Aisle (View All BBQ or Steak & Ribs)
		//				Shelf (ASDA Rewards or Burgers).
		// Skip links we don't care about, i.e. non-food departments.
		// Some pages will be paginated, need to look for the pagination controls and go to all the pages.
		// If we go to View All and paginate then we don't need to visit all the other links under a department.
		// Some sections end at Aisle with no shelves like Seafood & Fish.
		//if strings.Contains(shelfAisle.Text, "View All") {
		//	continue
		//}

		// Skip:
		// anything containing /event/ like /event/good-housekeeping-approved
		// /promotion/ like /promotion/2-for-4/
		// /cat/summer/
		// /cat/price-locked
		// /cat/eid-mubarak/
		// /cat/toiletries-beauty/
		// /cat/laundry-household/
		// /cat/baby-toddler-kids/
		// /cat/pet-food-accessories/

		// actually might be faster to allowlist all the ones we do care about

		log.Printf("Going to %s\n", shelfAisle.URL)
		driver.Get(shelfAisle.URL)

		// Wait for main product list to load.
		var mainProductDiv selenium.WebElement
		driver.Wait(func(wd selenium.WebDriver) (bool, error) {
			time.Sleep(50 * time.Millisecond)
			mainProductDiv, _ = driver.FindElement(selenium.ByClassName, mainProductClass)
			return mainProductDiv != nil, nil
		})

		// Get all anchor tags inside main product list.
		var productAnchors []selenium.WebElement
		driver.Wait(func(wd selenium.WebDriver) (bool, error) {
			time.Sleep(50 * time.Millisecond)
			productAnchors, _ = mainProductDiv.FindElements(selenium.ByTagName, aTag)
			return len(productAnchors) > 0, nil
		})

		// Iterate over anchors and get the urls.
		for _, a := range productAnchors {
			href, err := a.GetAttribute(hrefTag)
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

	fmt.Printf("%d product pages found", len(productPages))

	// Write all product page links to a json file.
	file, _ := json.MarshalIndent(productPages, "", " ")
	_ = os.WriteFile(productLinksFile, file, 0644)
}
