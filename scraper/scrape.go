package scraper

type Result struct {
	ConnectingFlight []
	HtmlBody string
}

type AirlineFlight struct {
	DepartureAirport string
	ArrivalAirport string
	FlightNumber string
	Airline string
	Class string
		
}


type Scraper interface {
	New(url string)
	scrape() []Result
}
