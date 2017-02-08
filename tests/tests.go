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
	die := func(msg string) {
		driver.DeleteSession()
		log.Fatalln(msg)
	}
	logTestResult := func(passed bool, err error, testDesc string) {
		doLog(statusText(passed && (err == nil)), "-", testDesc)
		if passed && err == nil {
			numPassed++
		} else {
			numFailed++
			if failFast {
				time.Sleep(5000 * time.Millisecond)
				die("Found first failing test, quitting")
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
	countCSSSelector := func(sel string) int {
		elements, xerr := driver.FindElements(goselenium.ByCSSSelector(sel))
		if xerr == nil {
			return len(elements)
		}
		return 0
	}
	cssSelectorExists := func(sel string) bool {
		count := countCSSSelector(sel)
		return (count != 0)
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

	welcomeCount := countCSSSelector(selectors.Welcome)
	logTestResult(welcomeCount == 0, nil, "Should not be welcoming anybody b/c nobody is logged in!")

	doLog("When trying to register, your site")

	err = submitForm(driver, selectors.LoginForm, users[0].loginFormData(), selectors.LoginFormSubmit)
	result = cssSelectorExists(selectors.Errors)
	logTestResult(result, err, "Should not allow unrecognized users to log in")

	badUsers := getBadUsers()
	for _, user := range badUsers {
		msg := "should not allow registration of a user with " + user.flaw
		err2 := registerUser(driver, testURL, user)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	err = registerUser(driver, testURL, users[0])
	if err == nil {
		result = cssSelectorExists(selectors.Welcome)
	}
	logTestResult(result, err, "Should welcome users that register with valid credentials")

	el, err := getEl(".logout")
	result = false
	if err == nil {
		el.Click()
		var response *goselenium.CurrentURLResponse
		response, err = driver.CurrentURL()
		if err == nil {
			var parsedURL *url.URL
			parsedURL, err = url.Parse(response.URL)
			if err == nil {
				result = parsedURL.Path == "/"
				if result {
					result = cssSelectorsExists(selectors.LoginForm, selectors.RegisterForm)
				}
			}
		}
	}
	logTestResult(result, err, "Should redirect users to '/' after logout")

	logout := func() {
		el, _ := getEl(".logout")
		result = false
		if err == nil {
			el.Click()
		}
	}

	// Register the other two users
	err = registerUser(driver, testURL, users[1])
	if err != nil {
		die("Error registering second user")
	}
	logout()
	err = registerUser(driver, testURL, users[2])
	if err != nil {
		die("Error registering third user")
	}
	logout()

	fmt.Println("A newly registered user")
	err = loginUser(driver, testURL, users[0])
	logTestResult(true, err, "Should be able to log in again")

	numTasks := countCSSSelector(selectors.Task)
	logTestResult(numTasks == 0, nil, "There should be no tasks at first")

	numTaskForms := countCSSSelector(selectors.TaskForm)
	logTestResult(numTaskForms == 1, nil, "There should a form for submitting tasks")

	badTasks := getBadTasks()
	for _, task := range badTasks {
		msg := "should not allow creation of a task with " + task.flaw
		err2 := submitTaskForm(driver, testURL, task)
		var count int
		if err2 == nil {
			count = countCSSSelector(selectors.Errors)
		}
		logTestResult(count == 1, err2, msg)
	}

	return numPassed, numFailed, err
}
