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

	task := randomTask(false)
	task.collaborator1 = users[1].email
	err = submitTaskForm(driver, testURL, task)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "There should be one task for user #1 after a valid task is submitted")

	task = randomTask(false)
	err = submitTaskForm(driver, testURL, task)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2, err, "There should be two tasks for user #1 after another is submitted")

	logout()
	_ = loginUser(driver, testURL, users[1])
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "User #2 should be able to log in and see the task that was shared with her")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskDelete) == 0, err, "She is not prompted to delete that task (she's not the owner)")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskCompleted) == 0, err, "The task is initially incomplete")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskComplete) == 0, err, "She is prompted to complete the task")
	el, err = getEl(selectors.TaskComplete)
	el.Click()
	logTestResult(countCSSSelector(selectors.TaskCompleted) == 1, err, "The task is marked as completed when she clicks the \"complete\" action")
	logout()

	_ = loginUser(driver, testURL, users[1])
	numCompleted := countCSSSelector(selectors.TaskCompleted)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2 && numCompleted == 1, err, "When user #0 logs in, one of the two tasks is marked as completed")
	logTestResult(countCSSSelector(selectors.TaskDelete) == 2, err, "User #0 is  prompted to delete both tasks (she's the owner)")
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	logTestResult(countCSSSelector(selectors.Task) == 1, err, "After deleting a task, only one is shown now")
	numTasks = countCSSSelector(selectors.Task)
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	logTestResult(numTasks == 1 && countCSSSelector(selectors.Task) == 0, err, "After deleting the next task, none are shown")

	return numPassed, numFailed, err
}
