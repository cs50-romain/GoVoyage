package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cs50-romain/GoVoyage/internal/flight"

	"github.com/gin-gonic/gin"
)

var flights []flight.Flight

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

	f := flight.New("FakeAirline", time.Now().Format(time.Kitchen), time.Now().Format(time.Kitchen), 300)
	fmt.Println(f)
	flights = append(flights, *f)
	//htmlElement := fmt.Sprintf(`<li>%s - Departure: %v, Arrival: %v, Price: %d`, f.Airline, f.Departure.Format(time.Kitchen), f.Arrival.Format(time.Kitchen), f.Price)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	log.Print("This is a test")

	router.GET("/", indexHandler)
	router.POST("/flights", flightsHandler)
	router.Run()
}
