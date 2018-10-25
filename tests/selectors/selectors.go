package selectors

const (
	// Every page ...
	Header          = "header"
	BootstrapHref   = "head link[rel=\"stylesheet\"][href*=\"bootstrap\"]"
	Footer          = "footer"
	FooterHomeLink  = "footer a[href=\"/\"]"
	FooterAboutLink = "footer a[href=\"/about\"]"
	Errors          = "ul[class=\"form-errors\"]"

	// home
	TeamLogo        = "* img[id=\"logo\"]"
	PageTitle       = "title:contains(\"event\")"
	EventList       = "ul li[class=\"event\"][id^=\"event-\"]"
	EventTime       = "li[class=\"event\"][id^=\"event-\"] time"
	EventDetailLink = "li[class=\"event\"][id^=\"event-\"] a[id=\"title\"][href*=\"/events/\"]" // how to specify title?
	NewEventLink    = "[id=\"new\"]"
	MobileResponse  = "div[id=\"mobile\"] :visible"
	DesktopResponse = "#response "

	// about
	Names     = "span[id$=\"-name\"]"    // how to get nickname programatically?
	Headshots = "img[id$=\"-headshot\"]" // ditto

	// new event
	NewEventForm          = "form[method=\"POST\"]"
	NewEventTitle         = "form input[name=\"title\"]"
	NewEventTitleLabel    = "form label[for=\"title\"]"
	NewEventImage         = "form input[name=\"image\"]"
	NewEventImageLabel    = "form label[for=\"image\"]"
	NewEventLocation      = "form input[name=\"location\"]"
	NewEventLocationLabel = "form label[for=\"location\"]"
	NewEventYear          = "form select[name=\"year\"]"
	NewEventYearOption    = "form select[name=\"year\"] option"
	NewEventYearLabel     = "form label[for=\"year\"]"
	NewEventMonth         = "form select[name=\"month\"]"
	NewEventMonthOption   = "form select[name=\"month\"] option"
	NewEventMonthLabel    = "form label[for=\"month\"]"
	NewEventDay           = "form select[name=\"day\"]"
	NewEventDayLabel      = "form label[for=\"day\"]"
	NewEventDayOption     = "form select[name=\"day\"] option"
	NewEventHour          = "form select[name=\"hour\"]"
	NewEventHourLabel     = "form label[for=\"hour\"]"
	NewEventHourOption    = "form select[name=\"hour\"] option"
	NewEventMinute        = "form select[name=\"minute\"]"
	NewEventMinuteLabel   = "form label[for=\"minute\"]"
	NewEventMinuteOption  = "form select[name=\"minute\"] option"
	NewEventSubmit        = "form [name=\"submit\"]"

	// event detail
	EventTitle        = "h1[id=\"title\"]"
	EventDate         = "span[id=\"date\"]"
	EventLocation     = "span[id=\"location\"]"
	EventImage        = "img[id=\"image\"]"
	EventAttendees    = "ul[id=\"attendees\"] li"
	RsvpEmail         = "form[method=\"POST\"] input[type=\"email\"][name=\"email\"]"
	RsvpEmailSubmit   = "form [name=\"submit\"]"
	EventDonationLink = "a[id=\"donate\"]"
)
