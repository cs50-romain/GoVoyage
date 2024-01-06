package flight

import (
	"cs50-romain/GoVoyage/internal/flightapi"
	"log"
)

type ResponseFlights struct {
	Airline			string
	DepartureAirport	string
	DepartureTime		string
	ArrivalAirport		string
	ArrivalTime		string
	Duration		int
	Price			int
}

type RequestedFlight struct {
	Airline		string
	Departure	string
	Arrival		string
	Price		int
}

func New(airline string, departure, arrival string, price int) *RequestedFlight {
	return &RequestedFlight{
		Airline: airline,
		Departure: departure,
		Arrival: arrival,
		Price: price,
	}
}

func GetFlights(departure_city, arrival_city, departure_date, return_date string) ([]ResponseFlights, error) {
	departure_airport_code := GetCityAirportCode(departure_city)
	arrival_airport_code := GetCityAirportCode(arrival_city)
	all_flights_api_response, err := flightapi.GetFlights(departure_airport_code, arrival_airport_code, departure_date, return_date)
	if err != nil {
		log.Printf("[ERROR] API JSON Response Error -> %s\n", err)
		return nil, err
	}

	// Iterate through those flights and append them to a ResponseFlights array
	var result_flights []ResponseFlights
	for _, bestflight := range all_flights_api_response.BestFlights {
		for _, flight := range bestflight.Flights {
			//fmt.Printf("Flight #%d: Departure Airport: %s, Arrival Airport: %s, Departure Time: %s, Arrival Time: %s; Flight Duration: %dmin, Price: $%d\n", idx, flight.DepartureAirport.Name, flight.ArrivalAirport.Name, flight.DepartureAirport.Time, flight.ArrivalAirport.Time, bestflight.TotalDuration, bestflight.Price)

			result_flights = append(result_flights, ResponseFlights{
				Airline: flight.Airline,
				DepartureAirport: flight.DepartureAirport.Name,
				DepartureTime: flight.DepartureAirport.Time,
				ArrivalAirport: flight.ArrivalAirport.Name,
				ArrivalTime: flight.ArrivalAirport.Time,
				Duration: bestflight.TotalDuration,
				Price: bestflight.Price,

			})
		}
	}
	return result_flights, nil
}

func GetCityAirportCode(city string) string {
	cityAirportMap := make(map[string]string)
	// Add major cities and their airport codes to the map
	cityAirportMap["New York"] = "JFK"
	cityAirportMap["Los Angeles"] = "LAX"
	cityAirportMap["Chicago"] = "ORD"
	cityAirportMap["London"] = "LHR"
	cityAirportMap["Tokyo"] = "HND"
	cityAirportMap["Paris"] = "CDG"
	cityAirportMap["Beijing"] = "PEK"
	cityAirportMap["Dubai"] = "DXB"
	cityAirportMap["Sydney"] = "SYD"
	cityAirportMap["Toronto"] = "YYZ"
	cityAirportMap["Hong Kong"] = "HKG"
	cityAirportMap["Mumbai"] = "BOM"
	cityAirportMap["Singapore"] = "SIN"
	cityAirportMap["Istanbul"] = "IST"
	cityAirportMap["San Francisco"] = "SFO"
	cityAirportMap["Dublin"] = "DUB"
	cityAirportMap["Frankfurt"] = "FRA"
	cityAirportMap["Mexico City"] = "MEX"
	cityAirportMap["Seoul"] = "ICN"
	cityAirportMap["Sao Paulo"] = "GRU"
	cityAirportMap["Moscow"] = "SVO"
	cityAirportMap["Bangkok"] = "BKK"
	cityAirportMap["Cairo"] = "CAI"
	cityAirportMap["Amsterdam"] = "AMS"
	cityAirportMap["Johannesburg"] = "JNB"
	cityAirportMap["Rome"] = "FCO"
	cityAirportMap["Madrid"] = "MAD"
	cityAirportMap["Berlin"] = "TXL"
	cityAirportMap["Vienna"] = "VIE"
	cityAirportMap["Athens"] = "ATH"
	cityAirportMap["Stockholm"] = "ARN"
	cityAirportMap["Lisbon"] = "LIS"
	cityAirportMap["New Delhi"] = "DEL"
	cityAirportMap["Oslo"] = "OSL"
	cityAirportMap["Warsaw"] = "WAW"
	cityAirportMap["Brussels"] = "BRU"
	cityAirportMap["Copenhagen"] = "CPH"
	cityAirportMap["Helsinki"] = "HEL"
	cityAirportMap["Buenos Aires"] = "EZE"
	cityAirportMap["Lima"] = "LIM"
	cityAirportMap["Bogota"] = "BOG"
	cityAirportMap["Caracas"] = "CCS"
	cityAirportMap["Lagos"] = "LOS"
	cityAirportMap["Nairobi"] = "NBO"
	cityAirportMap["Cape Town"] = "CPT"
	cityAirportMap["Montreal"] = "YUL"
	cityAirportMap["Vancouver"] = "YVR"
	cityAirportMap["Melbourne"] = "MEL"
	cityAirportMap["Auckland"] = "AKL"
	cityAirportMap["Miami"] = "MIA"
	cityAirportMap["Baltimore"] = "BWI"

	return cityAirportMap[city]
}
