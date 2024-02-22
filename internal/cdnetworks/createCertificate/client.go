package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
		port            = 443
	)

	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService("C:\\Users\\lunyi\\Downloads\\chromedriver-win64\\chromedriver.exe", 4444)
	//service, err := selenium.NewChromeDriverService("/root/chrome-linux64/chrome", 80)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
		return
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally
	capabilities := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(capabilities, "http://localhost:4444/wd/hub")
	if err != nil {
		fmt.Printf("Failed to open session: %v\n", err)
		return
	}
	//defer driver.Quit()

	// Perform automation tasks
	err = wd.Get("https://dash.cdnetworks.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %v\n", err)
		return
	}

	// Wait for the page to load
	time.Sleep(1 * time.Second)

	// Find the input fields and enter values
	accountInput, err := wd.FindElement(selenium.ByID, "js-accountinput")
	if err != nil {
		fmt.Println("Failed to find account input:", err)
		return
	}
	accountInput.SendKeys("raypro")

	pwdInput, err := wd.FindElement(selenium.ByID, "js-pwdinput")
	if err != nil {
		fmt.Println("Failed to find password input:", err)
		return
	}
	pwdInput.SendKeys("")

	// Find and click the submit button
	submitBtn, err := wd.FindElement(selenium.ByXPATH, "//button[@type='submit']")
	if err != nil {
		fmt.Println("Failed to find submit button:", err)
		return
	}
	submitBtn.Click()

	time.Sleep(1 * time.Second)

	err = wd.Get("https://portalcdn.cdnetworks.com/v2/index#/certificate/oneKey/list?code=cert_mylist&parentCode=cert_ca&productCode=certificatemanagement")
	if err != nil {
		fmt.Printf("Failed to load page: %v\n", err)
		return
	}

	// Wait for the page to load

	time.Sleep(20 * time.Second)
	fmt.Println("Press any key to quit...")
	var input string
	fmt.Scanln(&input)

}
