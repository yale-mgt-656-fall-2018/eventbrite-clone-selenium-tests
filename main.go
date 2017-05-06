package main

import (
	"net/url"
	"os"
	"strings"
	"log"

	todotests "github.com/yale-cpsc-213/social-todo-selenium-tests/tests"
)

// You should run this something like
//
// > social-todo-selenium-tests "http://localhost:4444/wd/hub" "http://localhost:8000"
//
// if you are using Selenium and
//
// > social-todo-selenium-tests "http://localhost:9515" "http://localhost:8000"
//
// if you are using "naked" chromedriver. Of course, the port
// will depend on how you are running it.
//
func main() {
	usage := "social-todo-selenium-tests SELENIUM_URL TEST_URL [-fast]"
	if len(os.Args) < 3 {
		log.Fatal(usage)
	}
	log.SetFlags(log.Lshortfile)
	failFast := false
	if len(os.Args) >= 4 && strings.Contains(os.Args[3], "fast") {
		failFast = true
	}
	todotests.RunForURL(os.Args[1], os.Args[2], failFast)
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err == nil {
		return true
	}
	return false
}
