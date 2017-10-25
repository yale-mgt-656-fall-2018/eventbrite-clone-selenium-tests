package tests

import (
	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

func fillForm(driver goselenium.WebDriver, formSelector string, data map[string]string) error {
	getEl := func(sel string) (goselenium.Element, error) {
		return driver.FindElement(goselenium.ByCSSSelector(sel))
	}
	_, err := getEl(formSelector)
	if err != nil {
		return err
	}
	for sel, val := range data {
		el, err := getEl(sel)
		if err != nil {
			return err
		}
		el.SendKeys(val)
	}
	return nil
}

func submitForm(driver goselenium.WebDriver, formSelector string, data map[string]string, submitSelector string) error {
	err := fillForm(driver, formSelector, data)
	if err != nil {
		return err
	}
	el, err := driver.FindElement(goselenium.ByCSSSelector(submitSelector))
	if err != nil {
		return err
	}
	el.Click()
	return nil
}

func fillRSVPForm(driver goselenium.WebDriver, testURL string, rsvp RSVP) error {
	err2 := loadHome(driver, testURL)
	if err2 != nil {
		return err2
	}
	err2 = submitForm(driver, selectors.RsvpEmail, rsvp.sendRSVP(), selectors.RsvpEmailSubmit)
	return err2
}

func fillEventForm(driver goselenium.WebDriver, testURL string, event Event) error {
	err2 := loadHome(driver, testURL)
	if err2 != nil {
		return err2
	}
	err2 = submitForm(driver, selectors.NewEventForm, event.createFormData(), selectors.NewEventSubmit)
	return err2
}

func loadHome(driver goselenium.WebDriver, targetURL string) error {
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
		return "✅  PASS"
	}
	return "❌  FAIL"
}
