package tests

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-mgt-656-fall-2018/eventbrite-clone-tests/tests/selectors"
)

// RunForURL - runs the test given a target URL
//
func RunForURL(seleniumURL string, testURL string, verbose bool, failFast bool, sleepDuration time.Duration) (int, int, error) {
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

	goselenium.SessionPageLoadTimeout(5)

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	return Run(driver, strings.TrimSuffix(testURL, "/"), verbose, failFast, sleepDuration)
}

type existanceTest struct {
	selector    string
	description string
}

func handleC9SplashPage(driver goselenium.WebDriver) (*goselenium.GoResponse, error) {
	// If this is a cloud9 URL, let's make sure our test is not
	// ruined by their "click-through" warning about app previews.
	// We do this by setting a cookie that makes it look like we
	// already visited.
	url, err := driver.CurrentURL()
	if err == nil && strings.Contains(strings.ToLower(url.URL), "c9users") {
		c := &goselenium.Cookie{
			Name:  "c9.live.user.click-through",
			Value: "ok",
		}
		driver.AddCookie(c)
	}
	return driver.Go(url.URL)
}

// Run - run all tests
//
func Run(driver goselenium.WebDriver, testURL string, verbose bool, failFast bool, sleepDuration time.Duration) (int, int, error) {

	// Track how many tests passed and failed
	numPassed := 0
	numFailed := 0

	// Log to the console, if we're in verbose mode
	doLog := func(args ...interface{}) {
		if verbose {
			fmt.Println(args...)
		}
	}

	// Log a test result to the console. Incrementing num passed/failed.
	logTestResult := func(passed bool, err error, testDesc string) {
		doLog(statusText(passed && (err == nil)), "-", testDesc)
		if passed && err == nil {
			numPassed++
		} else {
			numFailed++
			if failFast {
				time.Sleep(5000 * time.Millisecond)
				log.Fatalln("Found first failing test, quitting")
			}
		}
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
	checkGoodRsvps := func(eventNum int) {
		goodRsvps := getGoodRsvps()
		for _, rsvp := range goodRsvps {
			numOriginalRsvps := countCSSSelector(selectors.EventAttendees)
			msg := "should allow RSVP with " + rsvp.attribute
			err2 := fillRSVPForm(driver, testURL+"/events/"+fmt.Sprint(eventNum), rsvp)
			time.Sleep(sleepDuration)
			numNewRsvps := countCSSSelector(selectors.EventAttendees)
			result := (numNewRsvps == (numOriginalRsvps + 1))
			logTestResult(result, err2, msg)
		}
	}
	checkBadRsvps := func(eventNum int) {
		badRsvps := getBadRsvps()
		for _, rsvp := range badRsvps {
			msg := "should not allow RSVP with " + rsvp.flaw
			err2 := fillRSVPForm(driver, testURL+"/events/"+fmt.Sprint(eventNum), rsvp)
			time.Sleep(sleepDuration)
			result := cssSelectorExists(selectors.Errors)
			logTestResult(result, err2, msg)
		}
	}

	logExists := func(cssSelector string, description string) bool {
		result := cssSelectorExists(cssSelector)
		logTestResult(result, nil, description)
		return result
	}
	logAllExist := func(description string, cssSelectors ...string) bool {
		allExist := true
		for _, cssSelector := range cssSelectors {
			result := cssSelectorExists(cssSelector)
			if result == false {
				allExist = false
				break
			}
		}
		logTestResult(allExist, nil, description)
		return allExist
	}

	// Checks the structure of a randomly chosen event
	checkEvent := func(maxEventNum int) {
		eventNum := rand.Intn(3)
		doLog("\nEvent " + fmt.Sprint(eventNum) + " (randomly chosen):")
		time.Sleep(sleepDuration)

		_, _ = driver.Go(testURL + "/events/" + fmt.Sprint(eventNum))
		// logTestResult(true, err, "is reachable")

		existanceTests := []existanceTest{
			{selectors.BootstrapHref, "uses bootstrap"},
			{selectors.Header, "has a header"},
			{selectors.Footer, "has a footer"},
			{selectors.FooterAboutLink, "has a link to the about page in footer"},
			{selectors.FooterHomeLink, "has a link to the home page in footer"},
			{selectors.EventTitle, "has a title"},
			{selectors.EventDate, "has a date"},
			{selectors.EventLocation, "has a location"},
			{selectors.EventImage, "has an image"},
			{selectors.EventAttendees, "has a list of attendees"},
			{selectors.RsvpEmail, "has a form to RSVP"},
			{selectors.EventDonationLink, "has a link to donate"},
		}
		for _, t := range existanceTests {
			logExists(t.selector, t.description)
		}

		checkBadRsvps(eventNum)
		checkGoodRsvps(eventNum)
	}

	doLog("\nHome page:")
	_, err := driver.Go(testURL)
	if err == nil {
		_, err = handleC9SplashPage(driver)
	}

	homepageTests := []existanceTest{
		{selectors.PageTitle, "has a title that includes the string \"event\""},
		{selectors.BootstrapHref, "uses bootstrap"},
		{selectors.Header, "has a header"},
		{selectors.Footer, "has a footer"},
		{selectors.FooterAboutLink, "has a link to the about page in footer"},
		{selectors.FooterHomeLink, "has a link to the home page in footer"},
		{selectors.TeamLogo, "has your team logo"},
		{selectors.NewEventLink, "has a link to the new event page"},
		{selectors.EventList, "has a list of events"},
		{selectors.EventDetailLink, "each event links to a \"detail\" page"},
		{selectors.EventTime, "each event shows its time"},
	}
	for _, t := range homepageTests {
		logExists(t.selector, t.description)
	}

	// TODO: test mobile responsiveness. Likely need to inject some JS into
	// the DOM in order to accomplish this. Or, can we adjust the viewport
	// size via WebDriver?
	// doLog("\nMobile responsiveness:")
	// result = cssSelectorExists(selectors.DesktopResponse)
	// doLog(result)

	doLog("\nAbout page:")
	time.Sleep(sleepDuration)

	_, _ = driver.Go(testURL + "/about")
	// logTestResult(true, err, "should be reachable")

	result := cssSelectorExists(selectors.Names)
	logTestResult(result, nil, "has your names")

	result = cssSelectorExists(selectors.Headshots)
	logTestResult(result, nil, "shows your headshots (no necessarily real)")

	// Check the structure of one of the event pages
	checkEvent(3)

	doLog("\nNew event page:")
	time.Sleep(sleepDuration)

	_, err = driver.Go(testURL + "/events/new")
	// logTestResult(true, err, "is reachable")

	logExists(selectors.NewEventForm, "has a form for event submission")
	logAllExist(
		"the form has a title input field with label",
		selectors.NewEventTitle,
		selectors.NewEventTitleLabel,
	)
	logAllExist(
		"the form has a image input field with label",
		selectors.NewEventImage,
		selectors.NewEventImageLabel,
	)
	logAllExist(
		"the form has a location text field with label",
		selectors.NewEventLocation,
		selectors.NewEventLocationLabel,
	)
	logAllExist(
		"the form has a year dropdown field with label",
		selectors.NewEventYear,
		selectors.NewEventYearLabel,
	)
	logAllExist(
		"the form has a month dropdown field with label",
		selectors.NewEventMonth,
		selectors.NewEventMonthLabel,
	)
	logAllExist(
		"the form has a day dropdown field with label",
		selectors.NewEventDay,
		selectors.NewEventDayLabel,
	)
	logAllExist(
		"the form has a hour dropdown field with label",
		selectors.NewEventHour,
		selectors.NewEventHourLabel,
	)

	logAllExist(
		"the form has a minute dropdown field with label",
		selectors.NewEventMinute,
		selectors.NewEventMinuteLabel,
	)

	badEvents := getBadEvents()
	for _, event := range badEvents {
		msg := "should not allow user to submit event with " + event.flaw
		err2 := fillEventForm(driver, testURL+"/events/new", event)
		time.Sleep(sleepDuration)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	// client := http.Client{
	// 	Timeout: time.Second * 5,
	// }
	// response, err := client.PostForm(testURL+"/events/new", badEvents[0].getURLValues())
	// fmt.Println(response.
	// fmt.Println(err)
	// fmt.Println(response)
	// return 1, 1, nil

	apiTestData := createFormDataAPITest()
	msg := "should allow event creation with valid parameters, redirecting to the new event after creation"
	time.Sleep(sleepDuration)
	err2 := fillEventForm(driver, testURL+"/events/new", apiTestData)
	result = false
	if err2 == nil {
		src, err3 := driver.PageSource()
		if err3 == nil {
			// This is in imperfect test. We're testing to see if the
			// title of this new event occurs any where in the source
			// of the page returned. This is the broadest test we could
			// think of. Still, imperfect.
			result = strings.Contains(src.Source, apiTestData.Title)
		}
	}
	logTestResult(result, nil, msg)

	doLog("\nAPI:")
	time.Sleep(sleepDuration)
	success := testAPIResponse(testURL+"/api/events", func(ar apiResponse) bool {
		return true
	})
	logTestResult(success, nil, "should return valid JSON")

	success = testAPIResponse(testURL+"/api/events?search="+apiTestData.Title, func(ar apiResponse) bool {
		return len(ar.Events) == 1
	})
	logTestResult(success, nil, "should be searchable by event title")

	doLog(fmt.Sprintf("\n✅  Passed: %d", numPassed))
	doLog(fmt.Sprintf("\n❌  Failed: %d\n\n", numFailed))

	return numPassed, numFailed, err
}
