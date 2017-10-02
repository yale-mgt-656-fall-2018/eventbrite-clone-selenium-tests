# Eventbrite clone automated tests

These are the tests for the Eventbrite clone.

## Running the code

You can download binaries here:

* // NEED NEW BINARIES

You'll need to make the code executable (likely `chmod` or the Windows equivalent). The code runs the tests by automating a browser via the [WebDriver](https://www.w3.org/TR/webdriver/) protocol. When I completed the assignment, I used either [Selenium stand-alone server](http://www.seleniumhq.org/download/) or
[ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) as my interface to Chrome.

Once you have one of those, you can run either

```
selenium-server -port 4444
```

or

```
chromedriver --port 4444
```

Clearly, you can use whatever port you want. Then, you run the test code like this

```
eventbrite-clone-selenium-tests "http://localhost:4444/wd/hub" "http://localhost:8000"
```

if you are using selenium or

```
eventbrite-clone-selenium-tests "http://localhost:4444" "http://localhost:8000"
```

if you are using ChromeDriver (with which you don't need `/wd/hub`). That will run
the tests against your app running on port `8000`. You can, of course, point it anywhere,
including at my demo apps. In my experience, ChromeDriver is a little faster than Selenium.

## Expected output

When all the tests pass, you should see this.

```
Home page:
😎  PASS - looks 💯  with Bootstrap CSS 
😎  PASS - has a header
😎  PASS - has a footer
😎  PASS - footer links to home page
😎  PASS - footer links to about page
😎  PASS - has your team logo
😎  PASS - shows a list of events
😎  PASS - individual events link to details and show time
😎  PASS - has a link to the new event page

About page:
😎  PASS - layout is correct
😎  PASS - has your names
😎  PASS - shows your headshots

Event 0:
😎  PASS - uses bootstrap
😎  PASS - has a header
😎  PASS - has a footer
😎  PASS - has a link to the about page in footer
😎  PASS - has a link to the home page in footer
😎  PASS - has a title
😎  PASS - has a date
😎  PASS - has a location
😎  PASS - has an image
😎  PASS - has a list of attendees
😎  PASS - has a form to RSVP
😎  PASS - should not allow RSVP with invalid email
😎  PASS - should not allow RSVP with non-yale email
😎  PASS - should allow RSVP with normal yale email
😎  PASS - should allow RSVP with scrambled yale email
😎  PASS - should allow RSVP with another scrambled yale email

Event 1:
😎  PASS - uses bootstrap
😎  PASS - has a header
😎  PASS - has a footer
😎  PASS - has a link to the about page in footer
😎  PASS - has a link to the home page in footer
😎  PASS - has a title
😎  PASS - has a date
😎  PASS - has a location
😎  PASS - has an image
😎  PASS - has a list of attendees
😎  PASS - has a form to RSVP
😎  PASS - should not allow RSVP with invalid email
😎  PASS - should not allow RSVP with non-yale email
😎  PASS - should allow RSVP with normal yale email
😎  PASS - should allow RSVP with scrambled yale email
😎  PASS - should allow RSVP with another scrambled yale email

Event 2:
😎  PASS - uses bootstrap
😎  PASS - has a header
😎  PASS - has a footer
😎  PASS - has a link to the about page in footer
😎  PASS - has a link to the home page in footer
😎  PASS - has a title
😎  PASS - has a date
😎  PASS - has a location
😎  PASS - has an image
😎  PASS - has a list of attendees
😎  PASS - has a form to RSVP
😎  PASS - should not allow RSVP with invalid email
😎  PASS - should not allow RSVP with non-yale email
😎  PASS - should allow RSVP with normal yale email
😎  PASS - should allow RSVP with scrambled yale email
😎  PASS - should allow RSVP with another scrambled yale email

New event page:
😎  PASS - layout is correct
😎  PASS - has a form for event submission
😎  PASS - has a correctly labeled title field
😎  PASS - has a correctly labeled image field
😎  PASS - has a correctly labeled location field
😎  PASS - has a labeled year field with correct options
😎  PASS - has a labeled month field with correct options
😎  PASS - has a labeled day field with correct options
😎  PASS - has a labeled hour field with correct options
😎  PASS - has a labeled minute field with correct options
😎  PASS - should not allow event with no title
😎  PASS - should not allow event with too-long title
😎  PASS - should not allow event with no image
😎  PASS - should not allow event with bad image
😎  PASS - should not allow event with no location
😎  PASS - should not allow event with too-long location
😎  PASS - should allow event with legal parameters

API:
😎  PASS - should return valid JSON
😎  PASS - should be searchable

✅  Passed: 79
❌  Failed: 0
```

You can watch a video of me running the tests here: // add when finished.

## Tips

* Most of the tests are in the `tests.go` file.
* The CSS selectors we are using to interact with your app are in the `selectors.go` file. You'll need to structure your HTML and CSS such that the code passes.
* Use the `-fast` flag if you want to save some time.
* If you want to alter this code, e.g. slow down your browser, simply clone with repo into your `$GOPATH`. On my system, this code lives at `/Users/kljensen/go/src/github.com/yale-mgt-656/eventbrite-clone-selenium-tests`. Then run `go get` and `go build`. This was built with go1.8.3.
