package tests

import (
	"fmt"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
)

// Run - run all tests
//
func Run(driver goselenium.WebDriver, testURL string, verbose bool) (int, int, error) {

	doLog := func(args ...interface{}) {
		if verbose {
			fmt.Println(args...)
		}
	}

	// Navigate to the HackerNews website.
	_, err := driver.Go(testURL)
	if err != nil {
		doLog(err)
		return 0, 0, err
	}

	// Click the 'new' link at the top
	el, err := driver.FindElement(goselenium.ByCSSSelector("a[href='newest']"))
	if err != nil {
		fmt.Println(err)
		return 0, 0, err
	}

	// Click the link.
	_, err = el.Click()
	if err != nil {
		fmt.Println(err)
		return 0, 0, err
	}

	// Wait until the URL has changed with a timeout of 1 second and a check
	// interval of 10ms..
	newLink := "https://news.ycombinator.com/newest"
	ok := driver.Wait(goselenium.UntilURLIs(newLink), 1*time.Second, 10*time.Millisecond)
	if !ok {
		fmt.Println("Wait timed out :<")
		return 0, 0, err
	}

	// Woohoo! We have successfully navigated to a page.
	fmt.Println("Successfully navigated to URL " + newLink)
	return 0, 0, err

}
