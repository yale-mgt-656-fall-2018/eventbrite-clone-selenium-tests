package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

// RunForURL - runs the test given a target URL
//
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

	goselenium.SessionPageLoadTimeout(1)

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	return Run(driver, testURL, true, failFast, sleepDuration)
}

type existanceTest struct {
	selector    string
	description string
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
	checkGoodRsvps := func(eventNum int) int {
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
		return 1
	}
	checkBadRsvps := func(eventNum int) int {
		badRsvps := getBadRsvps()
		for _, rsvp := range badRsvps {
			msg := "should not allow RSVP with " + rsvp.flaw
			err2 := fillRSVPForm(driver, testURL+"/events/"+fmt.Sprint(eventNum), rsvp)
			time.Sleep(sleepDuration)
			result := cssSelectorExists(selectors.Errors)
			logTestResult(result, err2, msg)
		}
		return 1
	}

	logExists := func(selector string, description string) bool {
		result := cssSelectorExists(selector)
		logTestResult(result, nil, description)
		return result
	}

	checkEvent := func(eventNum int) int {
		driver.Go(testURL + "/events/" + fmt.Sprint(eventNum))

		doLog("\nEvent " + fmt.Sprint(eventNum) + ":")
		time.Sleep(sleepDuration)

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
		}
		for _, t := range existanceTests {
			logExists(t.selector, t.description)
		}

		checkBadRsvps(eventNum)

		checkGoodRsvps(eventNum)

		return 1
	}

	_, err := driver.Go(testURL)
	if err != nil {
		return 0, 0, err
	}

	time.Sleep(sleepDuration)

	doLog("\nHome page:")

	result := cssSelectorExists(selectors.BootstrapHref)
	logTestResult(result, nil, "looks 💯  with Bootstrap CSS ")

	result = cssSelectorExists(selectors.Header)
	logTestResult(result, nil, "has a header")
	result = cssSelectorExists(selectors.Footer)
	logTestResult(result, nil, "has a footer")

	result = cssSelectorExists(selectors.FooterHomeLink)
	logTestResult(result, nil, "footer links to home page")
	result = cssSelectorExists(selectors.FooterAboutLink)
	logTestResult(result, nil, "footer links to about page")

	result = cssSelectorExists(selectors.TeamLogo)
	logTestResult(result, nil, "has your team logo")

	result = cssSelectorExists(selectors.EventList)
	logTestResult(result, nil, "shows a list of events")

	linkResult := cssSelectorExists(selectors.EventDetailLink)
	timeResult := cssSelectorExists(selectors.EventTime)
	logTestResult(linkResult && timeResult, nil, "individual events link to details and show time")

	result = cssSelectorExists(selectors.NewEventLink)
	logTestResult(result, nil, "has a link to the new event page")

	// doLog("\nMobile responsiveness:")
	// result = cssSelectorExists(selectors.DesktopResponse)
	// doLog(result)

	_, err = driver.Go(testURL + "/about")
	if err != nil {
		return 0, 0, err
	}

	doLog("\nAbout page:")
	time.Sleep(sleepDuration)

	bootstrapResult := cssSelectorExists(selectors.BootstrapHref)
	headerResult := cssSelectorExists(selectors.Header)
	footerResult := cssSelectorExists(selectors.Footer)
	footerHomeLinkResult := cssSelectorExists(selectors.FooterHomeLink)
	footerAboutLinkResult := cssSelectorExists(selectors.FooterAboutLink)

	logTestResult(bootstrapResult && headerResult && footerResult && footerHomeLinkResult && footerAboutLinkResult, nil, "layout is correct")

	result = cssSelectorExists(selectors.Names)
	logTestResult(result, nil, "has your names")

	result = cssSelectorExists(selectors.Headshots)
	logTestResult(result, nil, "shows your headshots")

	checkEvent(0)
	checkEvent(1)
	checkEvent(2)

	_, err = driver.Go(testURL + "/events/new")
	if err != nil {
		return 0, 0, err
	}

	doLog("\nNew event page:")
	time.Sleep(sleepDuration)

	bootstrapResult = cssSelectorExists(selectors.BootstrapHref)
	headerResult = cssSelectorExists(selectors.Header)
	footerResult = cssSelectorExists(selectors.Footer)
	footerHomeLinkResult = cssSelectorExists(selectors.FooterHomeLink)
	footerAboutLinkResult = cssSelectorExists(selectors.FooterAboutLink)

	logTestResult(bootstrapResult && headerResult && footerResult && footerHomeLinkResult && footerAboutLinkResult, nil, "layout is correct")

	result = cssSelectorExists(selectors.NewEventForm)
	logTestResult(result, nil, "has a form for event submission")

	titleResult := cssSelectorExists(selectors.NewEventTitle)
	titleLabelResult := cssSelectorExists(selectors.NewEventTitleLabel)
	logTestResult(titleResult && titleLabelResult, nil, "has a correctly labeled title field")

	imageResult := cssSelectorExists(selectors.NewEventImage)
	imageLabelResult := cssSelectorExists(selectors.NewEventImageLabel)
	logTestResult(imageResult && imageLabelResult, nil, "has a correctly labeled image field")

	locationResult := cssSelectorExists(selectors.NewEventLocation)
	locationLabelResult := cssSelectorExists(selectors.NewEventLocationLabel)
	logTestResult(locationResult && locationLabelResult, nil, "has a correctly labeled location field")

	yearResult := cssSelectorExists(selectors.NewEventYear)
	yearLabelResult := cssSelectorExists(selectors.NewEventYearLabel)
	yearOptionResult := countCSSSelector(selectors.NewEventYearOption)
	logTestResult(yearResult && yearLabelResult && yearOptionResult == 2, nil, "has a labeled year field with correct options")

	monthResult := cssSelectorExists(selectors.NewEventMonth)
	monthLabelResult := cssSelectorExists(selectors.NewEventMonthLabel)
	monthOptionResult := countCSSSelector(selectors.NewEventMonthOption)
	logTestResult(monthResult && monthLabelResult && monthOptionResult == 12, nil, "has a labeled month field with correct options")

	dayResult := cssSelectorExists(selectors.NewEventDay)
	dayLabelResult := cssSelectorExists(selectors.NewEventDayLabel)
	dayOptionResult := countCSS
Selector(selectors.NewEventDayOption)
	logTestResult(dayResult && dayLabelResult && dayOptionResult == 31, nil, "has a labeled day field with correct options")

	hourResult := cssSelectorExists(selectors.NewEventHour)
	hourLabelResult := cssSelectorExists(selectors.NewEventHourLabel)
	hourOptionResult := countCSSSelector(selectors.NewEventHourOption)
	logTestResult(hourResult && hourLabelResult && hourOptionResult == 24, nil, "has a labeled hour field with correct options")

	minuteResult := cssSelectorExists(selectors.NewEventMinute)
	minuteLabelResult := cssSelectorExists(selectors.NewEventMinuteLabel)
	minuteOptionResult := countCSSSelector(selectors.NewEventMinuteOption)
	logTestResult(minuteResult && minuteLabelResult && minuteOptionResult == 2, nil, "has a labeled minute field with correct options")

	badEvents := getBadEvents()
	for _, event := range badEvents {
		msg := "should not allow event with " + event.flaw
		err2 := fillEventForm(driver, testURL+"/events/new", event)
		time.Sleep(sleepDuration)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	apiTestData := createFormDataAPITest()
	msg := "should allow event with legal parameters"
	err2 := fillEventForm(driver, testURL+"/events/new", apiTestData)
	time.Sleep(sleepDuration)
	if err2 == nil {
		result = cssSelectorExists(selectors.RsvpEmail)
		// this isn't checking for HTTP status codes
	}
	logTestResult(result, err2, msg)

	// how to check for correct options, not just count?

	// _, err = driver.Go(testURL + "/api/events")
	doLog("\nAPI:")
	// time.Sleep(sleepDuration)

	type EventJSON struct {
		ID        int      `json:"id"`
		Title     string   `json:"title"`
		Date      string   `json:"date"`
		Image     string   `json:"image"`
		Location  string   `json:"location"`
		Attending []string `json:"attending"`
	}

	type APIResponse struct {
		Events []EventJSON `json:"events"`
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	success := true

	req, reqErr := http.NewRequest(http.MethodGet, testURL+"/api/events", nil)
	if reqErr != nil {
		success = false
	}

	res, resErr := client.Do(req)
	if resErr != nil {
		success = false
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		success = false
	}

	allResponse := APIResponse{}
	jsonErr := json.Unmarshal(body, &allResponse)
	if jsonErr != nil {
		success = false
	}

	logTestResult(success, nil, "should return valid JSON")

	req, reqErr = http.NewRequest(http.MethodGet, testURL+"/api/events?search="+apiTestData.title, nil)
	if reqErr != nil {
		return 0, 0, err
	}

	res, resErr = client.Do(req)
	if resErr != nil {
		return 0, 0, err
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		return 0, 0, err
	}

	searchResponse := APIResponse{}
	jsonErr = json.Unmarshal(body, &searchResponse)
	if jsonErr != nil {
		return 0, 0, err
	}

	logTestResult((len(searchResponse.Events) == 1), nil, "should be searchable")
	// elements, err := driver.FindElements(goselenium.ByCSSSelector(selectors.MobileResponse))
	// if( elements[0].isDisplayed()){
	// 	doLog("woot")
	// }

	fmt.Printf("\n✅  Passed: %d", numPassed)
	fmt.Printf("\n❌  Failed: %d\n\n", numFailed)

	return numPassed, numFailed, err
}
