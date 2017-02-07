package tests

import goselenium "github.com/bunsenapp/go-selenium"

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
