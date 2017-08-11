package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	eventtests "github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests"
)

// You should run this something like
//
// > eventbrite-clone-selenium-tests "http://localhost:4444/wd/hub" "http://localhost:8000"
//
// if you are using Selenium and
//
// > eventbrite-clone-selenium-tests "http://localhost:9515" "http://localhost:8000"
//
// if you are using "naked" chromedriver. Of course, the port
// will depend on how you are running it.
//
func main() {
	usage := "eventbrite-clone-selenium-tests SELENIUM_URL TEST_URL [-fast]"
	if len(os.Args) < 3 {
		log.Fatal(usage)
	}
	log.SetFlags(log.Lshortfile)
	failFast := false
	if len(os.Args) >= 4 && strings.Contains(os.Args[3], "fast") {
		failFast = true
	}
	eventtests.RunForURL(os.Args[1], os.Args[2], failFast, 0)
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err == nil {
		return true
	}
	return false
}
