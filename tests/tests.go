package tests

import (
	"fmt"
	"log"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-cpsc-213/social-todo-selenium-tests/tests/selectors"
)

// Run - run all tests
//
func Run(driver goselenium.WebDriver, testURL string, verbose bool, failFast bool) (int, int, error) {
	numPassed := 0
	numFailed := 0
	doLog := func(args ...interface{}) {
		if verbose {
			fmt.Println(args...)
		}
	}
	logTestResult := func(passed bool, err error, testDesc string) {
		doLog(statusText(passed && (err == nil)), "-", testDesc)
		if passed && err == nil {
			numPassed++
		} else {
			numFailed++
			if failFast {
				log.Fatalln("Found first failing test, quitting")
			}
		}
	}

	users := []User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	doLog("Your site")

	// Navigate to the URL.
	_, err := driver.Go(testURL)
	logTestResult(true, err, "Should be up and running")

	getEl := func(sel string) (goselenium.Element, error) {
		return driver.FindElement(goselenium.ByCSSSelector(sel))
	}
	checkEl := func(sel string) error {
		_, err = getEl(sel)
		return err
	}

	err = checkEl(selectors.LoginForm)
	logTestResult(true, err, "Should have a login form")

	err = checkEl(selectors.RegisterForm)
	logTestResult(true, err, "Should have a registration form")

	err = submitForm(driver, selectors.LoginForm, users[0].loginFormData(), selectors.LoginFormSubmit)
	time.Sleep(2000 * time.Millisecond)

	return numPassed, numFailed, err
}

func statusText(pass bool) string {
	if pass {
		return "✅ PASS"
	}
	return "❌ FAIL"
}
