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
// add when finished
```

You can watch a video of me running the tests here: // add when finished.

## Tips

* Most of the tests are in the `tests.go` file.
* The CSS selectors we are using to interact with your app are in the `selectors.go` file. You'll need to structure your HTML and CSS such that the code passes.
* Use the `-fast` flag if you want to save some time.
* If you want to alter this code, e.g. slow down your browser, simply clone with repo into your `$GOPATH`. On my system, this code lives at `/Users/kljensen/go/src/github.com/yale-mgt-656/eventbrite-clone-selenium-tests`. Then run `go get` and `go build`. This was built with go1.8.3.
