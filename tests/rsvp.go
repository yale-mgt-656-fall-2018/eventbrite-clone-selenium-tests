package tests

import (
	// "math/rand"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

type RSVP struct {
	email     string
	flaw      string
	attribute string
}

func randomEmail(yaleEmail bool, scramble bool) RSVP {
	if yaleEmail == true {
		var name string
		var yale string
		if randomdata.Boolean() {
			name = randomdata.FirstName(randomdata.Male)
		} else {
			name = randomdata.FirstName(randomdata.Female)
		}
		name += "." + randomdata.LastName()
		if scramble == true {
			yale = ""
			y := randomdata.Boolean()
			a := randomdata.Boolean()
			l := randomdata.Boolean()
			e := randomdata.Boolean()
			e2 := randomdata.Boolean()
			d := randomdata.Boolean()
			u := randomdata.Boolean()
			if y == true {
				yale = yale + "Y"
			} else {
				yale = yale + "y"
			}
			if a == true {
				yale = yale + "A"
			} else {
				yale = yale + "a"
			}
			if l == true {
				yale = yale + "L"
			} else {
				yale = yale + "l"
			}
			if e == true {
				yale = yale + "E"
			} else {
				yale = yale + "e"
			}

			yale = yale + "."

			if e2 == true {
				yale = yale + "E"
			} else {
				yale = yale + "e"
			}
			if d == true {
				yale = yale + "D"
			} else {
				yale = yale + "d"
			}
			if u == true {
				yale = yale + "U"
			} else {
				yale = yale + "u"
			}
		} else {
			yale = "yale.edu"
		}
		return RSVP{
			email: name + "@" + yale,
		}
	}
	return RSVP{
		email: randomdata.Email(),
	}
}

func (r RSVP) sendRSVP() map[string]string {
	data := map[string]string{
		selectors.RsvpEmail: r.email,
	}
	return data
}

func getBadRsvps() []RSVP {
	var r RSVP
	var rsvps []RSVP

	r = randomEmail(false, false)
	r.email = "sdfsdfsd"
	r.flaw = "invalid email"
	rsvps = append(rsvps, r)

	r = randomEmail(false, false)
	r.flaw = "non-yale email"
	rsvps = append(rsvps, r)

	return rsvps
}

func getGoodRsvps() []RSVP {
	var r RSVP
	var rsvps []RSVP

	r = randomEmail(true, false)
	r.attribute = "normal yale email"
	rsvps = append(rsvps, r)

	r = randomEmail(true, true)
	r.attribute = "mixed case yale email"
	rsvps = append(rsvps, r)

	return rsvps
}
