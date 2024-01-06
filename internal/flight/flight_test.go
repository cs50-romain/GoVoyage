package flight

import (
	"testing"
)

func TestGetCityAirportCode(t *testing.T) {
	var tests = []struct {
		input string
		want string
	}{
		{"New York", "JFK"},
		{"Los Angeles", "LAX"},
		{"Tokyo", "HND"},
		{"Paris", "CDG"},
		{"Berlin", "TXL"},
		{"Hjkhjk", ""},
		{"", ""},
		{"Miami", "MIA"},
	}

	for _, test := range tests {
		if got := GetCityAirportCode(test.input); got != test.want {
		t.Errorf("got %s; want %s\n", got, test.want)		
		}
	}

}
