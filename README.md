# Social todo automated tests

These are the tests for the social todo.

## Running the code

You can download binaries here:

* [Windows](https://kljensen.s3.amazonaws.com/public/social-todo-selenium-tests/windows/social-todo-selenium-tests.exe)
* [Linux](https://kljensen.s3.amazonaws.com/public/social-todo-selenium-tests/linux/social-todo-selenium-tests)
* [Mac](https://kljensen.s3.amazonaws.com/public/social-todo-selenium-tests/mac/social-todo-selenium-tests)

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
social-todo-selenium-tests "http://localhost:4444/wd/hub" "http://localhost:8000"
```

if you are using selenium or

```
social-todo-selenium-tests "http://localhost:4444" "http://localhost:8000"
```

if you are using ChromeDriver (with which you don't need `/wd/hub`). That will run
the tests against your app running on port `8000`. You can, of course, point it anywhere,
including at my demo apps. In my experience, ChromeDriver is a little faster than Selenium.

## Expected output

When all the tests pass, you should see this.

```
When no user is logged in, your site
✅ PASS - should be up and running
✅ PASS - should have a login form
✅ PASS - should have a registration form
✅ PASS - should not be welcoming anybody b/c nobody is logged in!
When trying to register, your site
✅ PASS - should not allow unrecognized users to log in
✅ PASS - should not allow registration of a user with invalid email
✅ PASS - should not allow registration of a user with email that is more than 50 letters
✅ PASS - should not allow registration of a user with empty name
✅ PASS - should not allow registration of a user with name that is more than 50 letters
✅ PASS - should not allow registration of a user with empty password
✅ PASS - should not allow registration of a user with password that is more than 50 letters
✅ PASS - should welcome users that register with valid credentials
✅ PASS - should redirect users to '/' after logout
A newly registered user
✅ PASS - should be able to log in again
✅ PASS - should see no tasks at first
✅ PASS - should see a form for submitting tasks
✅ PASS - should not be able to create a task with no name
✅ PASS - should not be able to create a task with name is more than 500 letters
✅ PASS - should not be able to create a task with invalid email for collaborator1
✅ PASS - should not be able to create a task with invalid email for collaborator2
✅ PASS - should not be able to create a task with invalid email for collaborator3
✅ PASS - should see a task after a valid task is submitted
✅ PASS - should see two tasks after another is submitted
User #2, after logging in
✅ PASS - should be able to see the task that was shared with her by user #1
✅ PASS - should not be not prompted to delete that task (she's not the owner)
✅ PASS - should see the task as incomplete
✅ PASS - should be able to mark the the task as complete
✅ PASS - should see the task as complete after clicking the "toggle" action
User #1, after logging in
✅ PASS - should see one of the two tasks marked as completed
✅ PASS - should be able to mark that is incompleted when she clicks the "toggle" action
✅ PASS - should be prompted to delete both tasks (she's the owner)
✅ PASS - should only see one after deleting a task
✅ PASS - should see none after deleting two tasks
```

You can watch a video of me running the tests here: [https://youtu.be/FcqJj0U6a7M](https://youtu.be/FcqJj0U6a7M).
