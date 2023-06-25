package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

func main() {
	// Run Chrome browser
	service, err := selenium.NewChromeDriverService("chromedriver", 4444)
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
		//"--headless",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	driver.Get("https://groceries.asda.com")

	var accept selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		accept, _ = driver.FindElement(selenium.ByID, "onetrust-accept-btn-handler")
		return accept != nil, nil
	})

	//<menu data-auto-id="buttonMenuItem" type="menu" class="asda-btn asda-btn--clear asda-btn--fluid navigation-menu__item" aria-expanded="false" data-di-id="di-id-d052fc60-bb0f3692"><span class="navigation-menu__text">Groceries<span role="img" data-auto-id="" class="asda-icon asda-icon--charcoal asda-icon--tiny asda-icon--rotate0 navigation-menu__submenu-icon-chevron"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="744c7e0-e8489608" data-di-rand="1687714028831"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span></span></menu>
	//accept, err := driver.FindElement(selenium.ByID, "onetrust-accept-btn-handler")
	//if err != nil {
	//	log.Panic(err)
	//}
	//titleChangeCondition := func(wd WebDriver) (bool, error) {
	//	title, err := wd.Title()
	//	if err != nil {
	//		return false, err
	//	}
	//
	//	return title == newTitle, nil
	//}

	//time.Sleep(1 * time.Second)
	accept.Click()
	//<menu id="onetrust-accept-btn-handler" data-di-id="#onetrust-accept-btn-handler">I Accept</menu>

	var navMenus []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		navMenus, _ = driver.FindElements(selenium.ByClassName, "navigation-menu__item")
		return accept != nil, nil
	})

	//buttons, err := driver.FindElements(selenium.ByClassName, "navigation-menu__item")

	var groceries selenium.WebElement

	for _, menu := range navMenus {
		btnElement, _ := menu.FindElement(selenium.ByClassName, "navigation-menu__text")
		if btnElement != nil {
			if text, _ := btnElement.Text(); text == "Groceries" {
				groceries = btnElement
				break
			}
		}
	}

	//time.Sleep(1 * time.Second)

	err = groceries.Click()
	if err != nil {
		log.Panic(err)
	}

	var col1 selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		col1, _ = driver.FindElement(selenium.ByClassName, "h-nav__list--1-columns")
		return col1 != nil, nil
	})

	var itemsCol1 []selenium.WebElement
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		itemsCol1, _ = col1.FindElements(selenium.ByClassName, "h-nav__item-text")
		return itemsCol1 != nil, nil
	})

	for _, itemCol1 := range itemsCol1 {
		fmt.Println(itemCol1.Text())
	}
	for {
	}
	//	col1,err:=driver.FindElements(selenium.ByClassName,"h-nav__list--1-columns")
	//	if err!=nil{
	//		log.Panic(err)
	//	}
	//
	//fmt.Println(col1)

	//<span class="navigation-menu__text">
	//Groceries
	//<span role="img" data-auto-id="" class="asda-icon asda-icon--charcoal asda-icon--tiny asda-icon--rotate0 navigation-menu__submenu-icon-chevron">
	//<svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="744c7e0-e8489608" data-di-rand="1687714028831"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd">
	//</path>
	//</svg>
	//</span>
	//</span>

	//<ul class="h-nav__list h-nav__list--1-columns">
	//<li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--active h-nav__item-menu--bold" type="menu" data-di-id="di-id-c18a37a9-926820ff"> <span class="h-nav__item-text">Asda Rewards</span></menu></li>
	//<li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-dc53cce4-90d136d6"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="9cc92ad2-509091b" data-di-rand="1687716310913"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg>< span><span class="h-nav__item-text">Summer</span></menu></li>
	//<li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-56dc9c5e"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="9cc92ad2-5b71276d" data-di-rand="1687716310913"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Price Locked</span></menu></li>
	//<li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-a098cbff"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="11e126b-19f2a13e" data-di-rand="1687716310914"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Fruit, Veg &amp; Flowers</span></menu></li>
	//<li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-cbaa4ea7"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="11e126b-7fcb5c81" data-di-rand="1687716310914"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Meat, Poultry &amp; Fish</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-dc53cce4-741285d0"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="11e126b-e65d83a0" data-di-rand="1687716310914"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Bakery</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-ec90e6a0"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="b9a2750e-dc82a098" data-di-rand="1687716310915"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Chilled Food</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-98444d67"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="b9a2750e-248dcc35" data-di-rand="1687716310915"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Frozen Food</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-bca8635d-a1d2b1c6"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="b9a2750e-acd5b7dd" data-di-rand="1687716310916"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Food Cupboard</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-dc53cce4-5882ecf3"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="ab17dae0-d112ba4d" data-di-rand="1687716310916"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Drinks</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-e5b7a9fe"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="ab17dae0-1c88a72d" data-di-rand="1687716310916"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Beer, Wine &amp; Spirits</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-a69997f4"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="ab17dae0-1f5f37e2" data-di-rand="1687716310917"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Toiletries &amp; Beauty</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-540be24a"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="13abbd85-83f54839" data-di-rand="1687716310917"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Laundry &amp; Household</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-4e89e58c"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="13abbd85-47a77a01" data-di-rand="1687716310917"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Baby, Toddler &amp; Kids</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-140e7c79-174b7483"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="4bc8add3-6863fca" data-di-rand="1687716310918"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Pet Food &amp; Accessories</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-452c3c4f-d668619"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="4bc8add3-22cd49fc" data-di-rand="1687716310918"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Health &amp; Wellness</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-a38e657b"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="4bc8add3-b4588a77" data-di-rand="1687716310918"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Home &amp; Entertainment</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-78be9895"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="d62abd52-318f2b16" data-di-rand="1687716310919"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Free From...</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-4a464c-65708c06"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="f374cab6-1f1a74a5" data-di-rand="1687716310919"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Vegan &amp; Plant Based</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-452c3c4f-953704"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="f374cab6-28e5feaa" data-di-rand="1687716310919"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">World &amp; Local Food</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-41d743dc-47d92ef3"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="dfb51a7b-2a80751a" data-di-rand="1687716310920"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Organic</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-80c399b"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="c9dcffec-29a62cea" data-di-rand="1687716310920"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Eid Mubarak</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-df4b3b"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="c9dcffec-afc27342" data-di-rand="1687716310921"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Big Night In</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-bca8635d-61ce2c51"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="71609889-4d58afef" data-di-rand="1687716310921"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Extra Special</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-bca8635d-3742c1ce"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="71609889-285c9fca" data-di-rand="1687716310921"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Better For You</span></menu></li><li class="h-nav__item"><menu class="h-nav__item-menu h-nav__item-menu--has-children" type="menu" aria-expanded="true" data-di-id="di-id-c18a37a9-128c10b4"><span role="img" data-auto-id="" class="asda-icon asda-icon--dark-gray asda-icon--small asda-icon--rotate270 h-nav__item-icon"><svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" data-di-res-id="63d53767-6ae72147" data-di-rand="1687716310922"><path d="M16.46 19.339l9.193-9.193a.5.5 0 01.707 0l.707.708a.5.5 0 010 .707l-9.546 9.546a2 2 0 01-2.829 0L5.146 11.56a.5.5 0 010-.707l.708-.708a.5.5 0 01.707 0l9.192 9.193a.5.5 0 00.707 0z" fill="#3d3d3d" class="asda-icon__draw" fill-rule="evenodd"></path></svg></span><span class="h-nav__item-text">Price Match</span></menu></li></ul>

	//	fmt.Println(driver.PageSource())

	//for {
	//
	//}
}
