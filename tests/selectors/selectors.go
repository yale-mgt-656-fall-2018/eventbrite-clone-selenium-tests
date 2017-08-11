package selectors

const (
	// Every page ...
	Header          = "head"
	BootstrapHref   = "head link[rel=\"stylesheet\"][href*=\"bootstrap\"]"
	Footer          = "footer"
	FooterHomeLink  = "footer a[href=\"/\"]"
	FooterAboutLink = "footer a[href=\"/about\"]"
	Errors          = "ul li.form-errors"

	// home
	TeamLogo  = "img[id=\"logo\"]"
	EventList = "ul"
	Event     = "ul li.event"
	EventTime = "li.event time[datetime]"
	// EventTitle   = "li.event.h1" // how to specify title?
	NewEventLink = "[id=\"new\"]"

	// about
	Names     = "span[id*=\"name\"]"    // how to get nickname??
	Headshots = "img[id*=\"headshot\"]" // ditto

	// new event
	NewEventForm          = "form[method=\"POST\"]"
	NewEventTitle         = "form input[name=\"title\"]"
	NewEventTitleLabel    = "form label[for=\"title\"]"
	NewEventImage         = "form input[name=\"image\"]"
	NewEventImageLabel    = "form label[for=\"image\"]"
	NewEventLocation      = "form input[name=\"location\"]"
	NewEventLocationLabel = "form label[for=\"location\"]"
	NewEventYear          = "form input[name=\"year\"][type=\"select\"]"
	NewEventYearLabel     = "form label[for=\"year\"]"
	NewEventMonth         = "form input[name=\"month\"][type=\"select\"]"
	NewEventMonthLabel    = "form label[for=\"month\"]"
	NewEventDay           = "form input[name=\"day\"][type=\"select\"]"
	NewEventDayLabel      = "form label[for=\"day\"]"
	NewEventHour          = "form input[name=\"hour\"][type=\"select\"]"
	NewEventHourLabel     = "form label[for=\"hour\"]"
	NewEventMinute        = "form input[name=\"minute\"][type=\"select\"]"
	NewEventMinuteLabel   = "form label[for=\"minute\"]"
	NewEventSubmit        = "form input[name=\"submit\"][type=\"submit\"]"

	// event detail
	EventTitle     = "h1[id=\"title\"]"
	EventDate      = "span[id=\"date\"]"
	EventLocation  = "span[id=\"location\"]"
	EventImage     = "[id=\"image\"]"
	EventAttendees = "ul[id=\"attendees\"]"
	RsvpEmail      = "form[method=\"POST\"] input[type=\"email\"][id=\"email\"][name=\"email\"]"
)
