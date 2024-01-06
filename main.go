package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"

	"cs50-romain/GoVoyage/internal/flight"

	"github.com/gin-gonic/gin"
)

var departure_flights []flight.ResponseFlights
var return_flights []flight.ResponseFlights

func indexHandler(c *gin.Context) {
	fmt.Println(return_flights)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Flights": departure_flights,
		"ReturnFlights": return_flights,
	})
}

func flightsHandler(c *gin.Context) {
	origin := c.PostForm("origin")
	dest := c.PostForm("destination")
	start := c.PostForm("startDate")
	end := c.PostForm("endDate")

	log.Printf("Origin: %s, Destination: %s\nStart Date: %s, End Date: %s\n", origin, dest, start, end)
	
	// Need to get the flights based on use demand
	departure_flights = flight.GetFlights(origin, dest, start, end)
	return_flights = flight.GetFlights(dest, origin, end, end)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	log.Print("This is a test")

	router.GET("/", indexHandler)
	router.POST("/flights", flightsHandler)
	router.Run()
}


