package flight

type Flight struct {
	Airline		string
	Departure	string
	Arrival		string
	Price		int
}

func New(airline string, departure, arrival string, price int) *Flight {
	return &Flight{
		Airline: airline,
		Departure: departure,
		Arrival: arrival,
		Price: price,
	}
}
