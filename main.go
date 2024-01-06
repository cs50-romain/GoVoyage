package main

import (
	"fmt"
	"net/http"

	"cs50-romain/GoVoyage/internal/flight"

	"github.com/gin-gonic/gin"
)

var departure_flights []flight.ResponseFlights
var return_flights []flight.ResponseFlights
var err error
var message string

func indexHandler(c *gin.Context) {
	fmt.Println(return_flights)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Flights": departure_flights,
		"ReturnFlights": return_flights,
		"Message": message,
	})
	message = ""
}

func flightsHandler(c *gin.Context) {
	origin := c.PostForm("origin")
	dest := c.PostForm("destination")
	start := c.PostForm("startDate")
	end := c.PostForm("endDate")
	
	// Need to get the flights based on use demand
	departure_flights, err = flight.GetFlights(origin, dest, start, end)
	if err != nil {
		// Create a message for user
		message = err.Error()
	}
	return_flights, err = flight.GetFlights(dest, origin, end, end)
	if err != nil {
		message = err.Error()
	}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)
	router.POST("/flights", flightsHandler)
	router.Run()
}


