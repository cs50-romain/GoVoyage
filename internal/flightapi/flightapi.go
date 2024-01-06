package flightapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	g "github.com/serpapi/google-search-results-golang"
	//m "github.com/mitchellh/mapstructure"
)

type Flights struct {
	SearchMetadata struct {
		ID               string  `json:"id,omitempty"`
		Status           string  `json:"status,omitempty"`
		JSONEndpoint     string  `json:"json_endpoint,omitempty"`
		CreatedAt        string  `json:"created_at,omitempty"`
		ProcessedAt      string  `json:"processed_at,omitempty"`
		GoogleFlightsURL string  `json:"google_flights_url,omitempty"`
		RawHTMLFile      string  `json:"raw_html_file,omitempty"`
		PrettifyHTMLFile string  `json:"prettify_html_file,omitempty"`
		TotalTimeTaken   float64 `json:"total_time_taken,omitempty"`
	} `json:"search_metadata,omitempty"`
	SearchParameters struct {
		Engine       string `json:"engine,omitempty"`
		Hl           string `json:"hl,omitempty"`
		Gl           string `json:"gl,omitempty"`
		DepartureID  string `json:"departure_id,omitempty"`
		ArrivalID    string `json:"arrival_id,omitempty"`
		OutboundDate string `json:"outbound_date,omitempty"`
		ReturnDate   string `json:"return_date,omitempty"`
		Currency     string `json:"currency,omitempty"`
	} `json:"search_parameters,omitempty"`
	BestFlights []struct {
		Flights []struct {
			DepartureAirport struct {
				Name string `json:"name,omitempty"`
				ID   string `json:"id,omitempty"`
				Time string `json:"time,omitempty"`
			} `json:"departure_airport,omitempty"`
			ArrivalAirport struct {
				Name string `json:"name,omitempty"`
				ID   string `json:"id,omitempty"`
				Time string `json:"time,omitempty"`
			} `json:"arrival_airport,omitempty"`
			Duration         int      `json:"duration,omitempty"`
			Airplane         string   `json:"airplane,omitempty"`
			Airline          string   `json:"airline,omitempty"`
			AirlineLogo      string   `json:"airline_logo,omitempty"`
			TravelClass      string   `json:"travel_class,omitempty"`
			FlightNumber     string   `json:"flight_number,omitempty"`
			TicketAlsoSoldBy []string `json:"ticket_also_sold_by,omitempty"`
			Legroom          string   `json:"legroom,omitempty"`
			Extensions       []string `json:"extensions,omitempty"`
			Overnight        bool     `json:"overnight,omitempty"`
		} `json:"flights,omitempty"`
		TotalDuration   int `json:"total_duration,omitempty"`
		CarbonEmissions struct {
			ThisFlight          int `json:"this_flight,omitempty"`
			TypicalForThisRoute int `json:"typical_for_this_route,omitempty"`
			DifferencePercent   int `json:"difference_percent,omitempty"`
		} `json:"carbon_emissions,omitempty"`
		Price          int      `json:"price,omitempty"`
		Type           string   `json:"type,omitempty"`
		AirlineLogo    string   `json:"airline_logo,omitempty"`
		Extensions     []string `json:"extensions,omitempty"`
		DepartureToken string   `json:"departure_token,omitempty"`
	} `json:"best_flights,omitempty"`
	OtherFlights []struct {
		Flights []struct {
			DepartureAirport struct {
				Name string `json:"name,omitempty"`
				ID   string `json:"id,omitempty"`
				Time string `json:"time,omitempty"`
			} `json:"departure_airport,omitempty"`
			ArrivalAirport struct {
				Name string `json:"name,omitempty"`
				ID   string `json:"id,omitempty"`
				Time string `json:"time,omitempty"`
			} `json:"arrival_airport,omitempty"`
			Duration         int      `json:"duration,omitempty"`
			Airplane         string   `json:"airplane,omitempty"`
			Airline          string   `json:"airline,omitempty"`
			AirlineLogo      string   `json:"airline_logo,omitempty"`
			TravelClass      string   `json:"travel_class,omitempty"`
			FlightNumber     string   `json:"flight_number,omitempty"`
			TicketAlsoSoldBy []string `json:"ticket_also_sold_by,omitempty"`
			Legroom          string   `json:"legroom,omitempty"`
			Extensions       []string `json:"extensions,omitempty"`
			Overnight        bool     `json:"overnight,omitempty"`
		} `json:"flights,omitempty"`
		Layovers []struct {
			Duration int    `json:"duration,omitempty"`
			Name     string `json:"name,omitempty"`
			ID       string `json:"id,omitempty"`
		} `json:"layovers,omitempty"`
		TotalDuration   int `json:"total_duration,omitempty"`
		CarbonEmissions struct {
			ThisFlight          int `json:"this_flight,omitempty"`
			TypicalForThisRoute int `json:"typical_for_this_route,omitempty"`
			DifferencePercent   int `json:"difference_percent,omitempty"`
		} `json:"carbon_emissions,omitempty"`
		Price          int      `json:"price,omitempty"`
		Type           string   `json:"type,omitempty"`
		AirlineLogo    string   `json:"airline_logo,omitempty"`
		Extensions     []string `json:"extensions,omitempty"`
		DepartureToken string   `json:"departure_token,omitempty"`
	} `json:"other_flights,omitempty"`
	PriceInsights struct {
		LowestPrice       int     `json:"lowest_price,omitempty"`
		PriceLevel        string  `json:"price_level,omitempty"`
		TypicalPriceRange []int   `json:"typical_price_range,omitempty"`
		PriceHistory      [][]int `json:"price_history,omitempty"`
	} `json:"price_insights,omitempty"`
}
// API Stuff

func GetFlights(departure_id, arrival_id, departure_date, return_date string) *Flights {
	var flights Flights
	parameter := map[string]string{
	    "engine": "google_flights",
	    "departure_id": departure_id,
	    "arrival_id": arrival_id,
	    "outbound_date": departure_date,
	    "return_date": return_date,
	    "currency": "USD",
	    "hl": "en",
	    "api_key": "9c50a567049a317c3d0e7c6f50e377ef15c452d4ee3dd9f47d1eb0e8edc54293",
	}
	
	search := g.NewGoogleSearch(parameter, "9c50a567049a317c3d0e7c6f50e377ef15c452d4ee3dd9f47d1eb0e8edc54293")
	result, err := execute(&search, "/search", "json")
	if err != nil {
		log.Printf("[ERROR] Error getting results from Google Search -> %s\n", err)
	}

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("[ERROR] Error getting response body -> %s\n", err)
	}

	err = json.Unmarshal(body, &flights)
	if err != nil {
		log.Printf("[ERROR] Error Unmarshaling json response -> %s\n", err)
	}
	return &flights
}

// This is straight from serpapi's source code but for one minor change
// Inmstead of being a method of the Search type, it takes a parameter of type Search.
func execute(search *g.Search, path string, output string) (*http.Response, error) {
	query := url.Values{}
	if search.Parameter != nil {
		for k, v := range search.Parameter {
			query.Add(k, v)
		}
	}

	// api_key
	if len(search.ApiKey) != 0 {
		query.Add("api_key", search.ApiKey)
	}

	// engine
	if len(query.Get("engine")) == 0 {
		query.Set("engine", search.Engine)
	}

	// source programming language
	query.Add("source", "go")

	// set output
	query.Add("output", output)

	endpoint := "https://serpapi.com" + path + "?" + query.Encode()
	rsp, err := search.HttpSearch.Get(endpoint)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
