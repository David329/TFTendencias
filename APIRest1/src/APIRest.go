//Package main
package main

import (
	"log"
	"net/http"

	Routes "./Routes"

	"github.com/julienschmidt/httprouter"
)

func main() {

	//New Router
	router := httprouter.New()

	//Router-User
	router.GET("/users", Routes.GetAllUser)
	router.GET("/users/:id", Routes.GetUserByID)
	router.POST("/users", Routes.PostUser)
	router.PUT("/users/:id", Routes.PutUserByID)
	router.DELETE("/users/:id", Routes.DeleteUserByID)

	//Router-Flight
	router.GET("/flights", Routes.GetAllFlight)
	router.GET("/flights/:id", Routes.GetFlightByID)
	router.POST("/flights", Routes.PostFlight)
	router.PUT("/flights/:id", Routes.PutFlightByID)
	router.DELETE("/flights/:id", Routes.DeleteFlightByID)

	//Router-Booking
	router.GET("/bookings", Routes.GetAllBooking)
	router.GET("/bookings/:id", Routes.GetBookingByID)
	router.POST("/bookings", Routes.PostBooking)
	router.PUT("/bookings/:id", Routes.PutBookingByID)
	router.DELETE("/bookings/:id", Routes.DeleteBookingByID)

	//Escuchando el servidor
	log.Print("Escuchando en 127.0.0.1:8100...")
	http.ListenAndServe(":8100", router)
}
