package tests

import (
	"fmt"
	"log"
	"net/url"
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
				time.Sleep(5000 * time.Millisecond)
				driver.DeleteSession()
				log.Fatalln("Found first failing test, quitting")
			}
		}
	}

	users := []User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	doLog("When no user is logged in, your site")

	getEl := func(sel string) (goselenium.Element, error) {
		return driver.FindElement(goselenium.ByCSSSelector(sel))
	}
	cssSelectorExists := func(sel string) bool {
		_, xerr := getEl(sel)
		return (xerr == nil)
	}
	cssSelectorsExists := func(sels ...string) bool {
		for _, sel := range sels {
			if cssSelectorExists(sel) == false {
				return false
			}
		}
		return true
	}

	// Navigate to the URL.
	_, err := driver.Go(testURL)
	logTestResult(true, err, "Should be up and running")

	result := cssSelectorExists(selectors.LoginForm)
	logTestResult(result, nil, "Should have a login form")

	result = cssSelectorExists(selectors.RegisterForm)
	logTestResult(result, nil, "Should have a registration form")

	doLog("When trying to register, your site")

	err = submitForm(driver, selectors.LoginForm, users[0].loginFormData(), selectors.LoginFormSubmit)
	result = cssSelectorExists(selectors.Errors)
	logTestResult(result, err, "Should not allow unrecognized users to log in")

	badUsers := getBadUsers()
	for _, user := range badUsers {
		msg := "should not allow registration of a user with " + user.description
		err2 := registerUser(driver, testURL, user)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	err = registerUser(driver, testURL, users[0])
	if err == nil {
		result = cssSelectorExists(selectors.Welcome) && !cssSelectorExists(selectors.Errors)
	}
	logTestResult(result, err, "Should welcome users that register with valid credentials")

	el, _ := getEl(".logout")
	result = false
	if err == nil {
		el.Click()
		response, err := driver.CurrentURL()
		if err == nil {
			parsedURL, err := url.Parse(response.URL)
			if err == nil {
				result = parsedURL.Path == "/"
				if result {
					result = cssSelectorsExists(selectors.LoginForm, selectors.RegisterForm)
				}
			}
		}
	}
	logTestResult(result, err, "Should redirect users to '/' after logout")

	time.Sleep(2000 * time.Millisecond)
	return numPassed, numFailed, err
}

func registerUser(driver goselenium.WebDriver, testURL string, user User) error {
	err2 := loadLogin(driver, testURL)
	if err2 != nil {
		return err2
	}
	err2 = submitForm(driver, selectors.RegisterForm, user.registerFormData(), selectors.RegisterFormSubmit)
	return err2
}

func loadLogin(driver goselenium.WebDriver, targetURL string) error {
	const script = `
    var forms = document.getElementsByTagName('form');
    for(var i=0; i<forms.length; i+=1){
      var f = forms[i];
      f.setAttribute('novalidate', true);
    }
  `
	_, err1 := driver.Go(targetURL)
	if err1 != nil {
		return err1
	}
	_, err2 := driver.ExecuteScript(script)
	return err2
}

func statusText(pass bool) string {
	if pass {
		return "✅ PASS"
	}
	return "❌ FAIL"
}
