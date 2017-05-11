package tests

import (
	"fmt"
	"log"
	"net/url"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-cpsc-213/social-todo-selenium-tests/tests/selectors"
)

func RunForURL(seleniumURL string, testURL string, failFast bool, sleepDuration time.Duration) (int, int, error) {
	// Create capabilities, driver etc.
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.ChromeBrowser())

	driver, err := goselenium.NewSeleniumWebDriver(seleniumURL, capabilities)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	_, err = driver.CreateSession()
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	return Run(driver, testURL, true, failFast, sleepDuration)
}

// Run - run all tests
//
func Run(driver goselenium.WebDriver, testURL string, verbose bool, failFast bool, sleepDuration time.Duration) (int, int, error) {
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
	logTestResult(true, err, "should be up and running")

	time.Sleep(sleepDuration)
	result := cssSelectorExists(selectors.LoginForm)
	logTestResult(result, nil, "should have a login form")

	result = cssSelectorExists(selectors.RegisterForm)
	logTestResult(result, nil, "should have a registration form")

	welcomeCount := countCSSSelector(selectors.Welcome)
	logTestResult(welcomeCount == 0, nil, "should not be welcoming anybody b/c nobody is logged in!")

	doLog("When trying to register, your site")

	err = submitForm(driver, selectors.LoginForm, users[0].loginFormData(), selectors.LoginFormSubmit)
	time.Sleep(sleepDuration)
	result = cssSelectorExists(selectors.Errors)
	logTestResult(result, err, "should not allow unrecognized users to log in")

	badUsers := getBadUsers()
	for _, user := range badUsers {
		msg := "should not allow registration of a user with " + user.flaw
		err2 := registerUser(driver, testURL, user)
		time.Sleep(sleepDuration)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	err = registerUser(driver, testURL, users[0])
	if err == nil {
		time.Sleep(sleepDuration)
		result = cssSelectorExists(selectors.Welcome)
	}
	logTestResult(result, err, "should welcome users that register with valid credentials")

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
					time.Sleep(sleepDuration)
					result = cssSelectorsExists(selectors.LoginForm, selectors.RegisterForm)
				}
			}
		}
	}
	logTestResult(result, err, "should redirect users to '/' after logout")

	logout := func() {
		element, _ := getEl(".logout")
		result = false
		if err == nil {
			element.Click()
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
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.Welcome) == 1, err, "should be able to log in again")

	numTasks := countCSSSelector(selectors.Task)
	logTestResult(numTasks == 0, nil, "should see no tasks at first")

	numTaskForms := countCSSSelector(selectors.TaskForm)
	logTestResult(numTaskForms == 1, nil, "should see a form for submitting tasks")

	badTasks := getBadTasks()
	for _, task := range badTasks {
		msg := "should not be able to create a task with " + task.flaw
		err2 := submitTaskForm(driver, testURL, task)
		var count int
		if err2 == nil {
			time.Sleep(sleepDuration)
			count = countCSSSelector(selectors.Errors)
		}
		logTestResult(count == 1, err2, msg)
	}

	task := randomTask(false)
	task.collaborator1 = users[1].email
	err = submitTaskForm(driver, testURL, task)
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "should see a task after a valid task is submitted")

	task = randomTask(false)
	err = submitTaskForm(driver, testURL, task)
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2, err, "should see two tasks after another is submitted")
	time.Sleep(3000 * time.Millisecond)

	logout()
	fmt.Println("User #2, after logging in")
	_ = loginUser(driver, testURL, users[1])
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "should be able to see the task that was shared with her by user #1")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskDelete) == 0, err, "should not be not prompted to delete that task (she's not the owner)")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskCompleted) == 0, err, "should see the task as incomplete")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskToggle) == 1, err, "should be able to mark the the task as complete")
	el, err = getEl(selectors.TaskToggle)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.TaskCompleted) == 1, err, "should see the task as complete after clicking the \"toggle\" action")
	logout()

	_ = loginUser(driver, testURL, users[0])
	fmt.Println("User #1, after logging in")
	time.Sleep(sleepDuration)
	numCompleted := countCSSSelector(selectors.TaskCompleted)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2 && numCompleted == 1, err, "should see one of the two tasks marked as completed")
	el, err = getEl(selectors.TaskToggle)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.TaskCompleted) == 0, err, "should be able to mark that is incompleted when she clicks the \"toggle\" action")
	logTestResult(countCSSSelector(selectors.TaskDelete) == 2, err, "should be prompted to delete both tasks (she's the owner)")
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.Task) == 1, err, "should only see one after deleting a task")
	numTasks = countCSSSelector(selectors.Task)
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(numTasks == 1 && countCSSSelector(selectors.Task) == 0, err, "should see none after deleting two tasks")

	return numPassed, numFailed, err
}
