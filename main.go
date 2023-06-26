package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const buttonTag = "button"
const navItemIconClass = "h-nav__item-icon"
const url = "https://groceries.asda.com"
const driverBin = "chromedriver"
const port = 4444
const acceptButtonID = "onetrust-accept-btn-handler"
const navMenuClass = "navigation-menu__item"
const navMenuTextClass = "navigation-menu__text"
const groceriesText = "Groceries"

func main() {
	// Run Chrome browser
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

	log.Printf("Going to %s...\n", url)
	driver.Get(url)

	var accept selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(50 * time.Millisecond)
		accept, _ = driver.FindElement(selenium.ByID, acceptButtonID)
		return accept != nil, nil
	})

	log.Println("Clicking Accept...")
	accept.Click()

	var navMenus []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(50 * time.Millisecond)
		navMenus, _ = driver.FindElements(selenium.ByClassName, navMenuClass)
		return accept != nil, nil
	})

	var groceries selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		for _, menu := range navMenus {
			time.Sleep(50 * time.Millisecond)
			btnElement, _ := menu.FindElement(selenium.ByClassName, navMenuTextClass)
			if btnElement != nil {
				if text, _ := btnElement.Text(); text == groceriesText {
					groceries = btnElement
					return true, nil
				}
			}
		}

		return false, nil
	})

	log.Println("Clicking Groceries...")
	err = groceries.Click()
	if err != nil {
		log.Panic(err)
	}

	aisles := []selenium.WebElement{}

	aisles = openMenu(driver, 0, aisles)

	log.Println("Finished")
}

// openMenu opens all items on a menu. Called recursively until we run out of submenus.
// The first menu under Groceries opens a <ul> with class h-nav__list--1-columns.
// This can be considered menu level 1. It contains an <li> for each menu item.
// Hovering over a Groceries menu item opens a submenu that is menu level 2. The submenu is displayed
// by showing two <ul> elements with class h-nav__list--2-columns.
// The first is the Groceries menu and the second is the submenu.
// Each new submenu shows a new set of <ul> elements where the
// number (n) in the h-nav__list--n-columns class matches the menu level.
// The number of <ul> elements matches the number in h-nav__list--n-columns.
func openMenu(driver selenium.WebDriver, menuLevel int, aisles []selenium.WebElement) []selenium.WebElement {
	columnClass := fmt.Sprintf("h-nav__list--%d-columns", menuLevel+1)

	var navCol []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(50 * time.Millisecond)
		navCol, _ = driver.FindElements(selenium.ByClassName, columnClass)
		return navCol != nil, nil
	})

	if len(navCol) == 0 {
		return aisles
	}

	var allButtons []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(50 * time.Millisecond)
		allButtons, _ = navCol[menuLevel].FindElements(selenium.ByTagName, buttonTag)
		return allButtons != nil, nil
	})

	aisleButtons := []selenium.WebElement{}
	subMenuButtons := []selenium.WebElement{}
	for _, btn := range allButtons {
		hasSubMenuIcon, _ := btn.FindElements(selenium.ByClassName, navItemIconClass)
		if len(hasSubMenuIcon) == 0 {
			aisleButtons = append(aisleButtons, btn)
			continue
		}

		subMenuButtons = append(subMenuButtons, btn)
	}

	for _, subMenuButton := range subMenuButtons {
		err := subMenuButton.MoveTo(10, 10)
		if err != nil {
			log.Print("couldn't move to submenu button, skipping")
			continue
		}

		time.Sleep(200 * time.Millisecond)
		txt, _ := subMenuButton.Text()
		log.Printf("Opening menu %s\n", txt)
		aisles = openMenu(driver, menuLevel+1, aisles)
	}

	aisles = append(aisles, aisleButtons...)

	log.Printf("Aisle links %d\n", len(aisles))

	return aisles
}
