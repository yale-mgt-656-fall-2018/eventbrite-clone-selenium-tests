package tests

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

type Event struct {
	title    string
	location string
	image    string
	year     string
	month    string
	day      string
	hour     string
	minute   string
}

func randomEvent() Event {
	e := Event{
		title:    randomdata.Adjective() + " " + randomdata.Noun(),
		location: randomdata.Address(),
		image:    "idk.png", // ???
		year:     randomdata.Number(2015, 2016),
		month:    randomdata.Month(),
		day:      randomdata.Number(28),
		hour:     randomdata.Number(23),
		minute:   30,
	}
	return e
}

func (e Event) createFormData() map[string]string {
	data := map[string]string{
		selectors.NewEventTitle:    e.title,
		selectors.NewEventImage:    e.image,
		selectors.NewEventLocation: e.location,
		selectors.NewEventYear:     e.year,
		selectors.NewEventMonth:    e.month,
		selectors.NewEventDay:      e.day,
		selectors.NewEventHour:     e.hour,
		selectors.NewEventMinute:   e.minute,
	}
	return data
}

func getBadEvents() []Event {
	var e Event
	var events []Event

	e = randomEvent()
	e.flaw = "no title"
	e.name = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no image"
	e.image = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no location"
	e.location = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no year"
	e.year = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no month"
	e.month = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no day"
	e.day = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no hour"
	e.hour = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no minute"
	e.minute = ""
	events = append(events, e)

	// any other flaws?

	return events
}
