package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"

	"cs50-romain/GoVoyage/internal/flight"

	"github.com/gin-gonic/gin"
)

var flights []flight.ResponseFlights

func indexHandler(c *gin.Context) {
	fmt.Println(flights)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Flights": flights,
	})
}

func flightsHandler(c *gin.Context) {
	origin := c.PostForm("origin")
	dest := c.PostForm("destination")
	start := c.PostForm("startDate")
	end := c.PostForm("endDate")

	log.Printf("Origin: %s, Destination: %s\nStart Date: %s, End Date: %s\n", origin, dest, start, end)
	
	// Need to get the flights based on use demand
	flights = flight.GetFlights(origin, dest, start, end)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	log.Print("This is a test")

	router.GET("/", indexHandler)
	router.POST("/flights", flightsHandler)
	router.Run()
}


