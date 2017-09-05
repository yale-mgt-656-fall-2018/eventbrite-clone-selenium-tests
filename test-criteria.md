Here's what we're testing your app for:

All pages:
* Every page should look good. That means including the Bootstrap CSS stylesheet in the `head` of the page -- we'll be looking for a link with an `href` that contains `bootstrap`.
* Every page should have a `header`.
* Every page should have a `footer` with links to the home page and about page.
* Grading note: we'll test each of these individually for the home page, then test them as a group for subsequent pages -- if any one of them fails, all of them will. Also, we're defining 'every page' as `/`, `/about`, `/events/new`, `/events/0`, `/events/1`, and `/events/2`.

Home page:
* The home page should have your logo -- we'll look for an `img` tag with the id `logo`.
* The home page should also have a link to the new event page with the id `new`.
* Finally, the home page should have a list of events (a `ul` list, to be specific). Each event should have the class `event` and the id `event-x`, where `x` is the event's id number. Each event's title should also link to its detail page with an `a` element with the `id` 'title', and the event's `li` should contain a `time` tag with the event's date and time.

About page:
* The about page should have all team members' names in a `span` with the id `class-nickname-name`, where `class-nickname` is, well, your class nickname (check your profile on the course website if you're confused). Don't forget the dashes. Additionally, we'll look for a picture of each team member with the id `class-nickname-headshot`. (Note that this picture doesn't need to actually be of the member -- we're not using computer vision to make sure you have decent headshots or anything -- but you do need an image for each member.)

New event page:
* The new event page should have a form that POSTs its data. It needs to have inputs for `title`, `image`, `location`, `year`, `month`, `day`, `hour`, and `minute` as follows (make sure that you match these names exactly, otherwise the tests will fail):
    * Each form element should have a `name` attribute set to its name (`title`, etc.).
    * Each form element should be labeled with a `label` element with a `for` attribute set to the form element's name.
    * `year`, `month`, `day`, `hour`, and `minute` should all be dropdown menus. `year` should only have 2017 and 2018 as options, `month` should have all of the months' names (not numbers), `day` should have 1-31 (we're not checking for valid month/day combinations, though that would be a great extension), `hour` should be 0-23, and `minute` should only have 00 and 30 as options.
* Errors for new event creation are:
    * No title, image, or location
    * Title or location longer than 50 characters
    * Image with a file extension other than `.png` or `.gif`
    * Image that doesn't exist online (you should be testing this!)
* You'll want to show an error message if any of the conditions above are met. We'll be looking for them in a `ul` with the class `form-errors`.

Event detail pages:
* Event detail pages should have the event's title in a `h1` element with the id `title`, as well as the date and location in `span` elements with the ids `date` and `location`, respectively. The event's image should have the id `image`.
* Event detail pages should also have a `ul` list of the attendees with the id `attendees`. Each attendee (represented by their email) should be in a list element with the class `list-group-item`. (Note that this list might not show up if you don't have any attendees yet, so for grading purposes it's probably a good idea to RSVP to your own events. Isn't that good form, anyway?)
* Finally, event detail pages should have a POST form to RSVP. The input in this form should have name, id, and type all set to `email`.
    * This form should validate the email that you give it to make sure that it's a valid yale.edu email. However, it should accept any weird capitalizations as long as they're Yale emails, so something like 'kYle.JeNseN@yAlE.eDu' is acceptable (yes, this is just like that SpongeBob meme).

API:
* You should have an API running at `/api/events` that returns a valid JSON of every event that your app knows about. It should look like this:
```
    {events:
        [
            {
                id: 0,
                title: 'Hello world',
                location: 'New Haven',
                attendees: [
                    'jacob.bendicksen@yale.edu',
                ],
                image: 'http://yaleherald.com/wp-content/uploads/2013/02/Screen-Shot-2013-02-27-at-11.51.10-AM.png',
                time: '9/19/17 23:30'
            },
            <more events here, maybe>...
        ]
    }
```
* Grading note: the JSON reader that we're using for the tests is pretty finicky, so make sure that your API responses are structured exactly like this.
* If a `search` parameter is attached to the API request (something like `/api/events?search=hello`), you should only return events with the search term in the title. If there aren't any, return a JSON that looks like:
    * `{events: []}`

Note that while this document is our best effort at translating the tests into plain English, the grading code is the final source of truth. It's available to you to look through if you'd like (it's written in Go, which shouldn't be too hard to understand if you're feeling okay about JavaScript), and at the end of the day, your grade will be determined by how many of the tests you're passing in the grading code, rather than their descriptions here. We're happy to walk you through the grading code in office hours if you're confused!
