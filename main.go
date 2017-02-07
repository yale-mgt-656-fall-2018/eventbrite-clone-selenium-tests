package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	todotests "github.com/yale-cpsc-213/social-todo-selenium-tests/tests"

	goselenium "github.com/bunsenapp/go-selenium"
)

// You should run this something like
//
// > todotest "http://localhost:4444/wd/hub" "http://localhost:8000"
//
// if you are using Selenium and
//
// > todotest "http://localhost:9515" "http://localhost:8000"
//
// if you are using "naked" chromedriver. Of course, the port
// will depend on how you are running it.
//
func main() {
	usage := "todotest SELENIUM_URL TEST_URL"
	if len(os.Args) != 3 {
		log.Fatal(usage)
	}
	doRun(os.Args[1], os.Args[2])
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err == nil {
		return true
	}
	return false
}

func doRun(seleniumURL string, testURL string) {
	// Create capabilities, driver etc.
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.ChromeBrowser())

	driver, err := goselenium.NewSeleniumWebDriver(seleniumURL, capabilities)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	todotests.Run(driver, testURL, true)
}
