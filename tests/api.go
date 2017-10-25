package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type eventJSON struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Date      string   `json:"date"`
	Image     string   `json:"image"`
	Location  string   `json:"location"`
	Attending []string `json:"attending"`
}

type apiResponse struct {
	Events []eventJSON `json:"events"`
}

type apiResponseChecker func(apiResponse) bool

// Check to see if an the JSON API of an app works. Takes a
// function that accepts an `apiResponse` struct and returns
// a boolean. Returns false if that function returns false or
// if we encounter an error while trying to get & parse the
// API response.
//
func testAPIResponse(url string, checker apiResponseChecker) bool {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	success := false
	if reqErr == nil {
		res, resErr := client.Do(req)
		if resErr == nil {
			body, readErr := ioutil.ReadAll(res.Body)
			if readErr == nil {
				ar := apiResponse{}
				jsonErr := json.Unmarshal(body, &ar)
				if jsonErr == nil {
					success = checker(ar)
				}
			}
		}
	}
	return success
}
