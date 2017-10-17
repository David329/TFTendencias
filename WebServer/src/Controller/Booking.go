package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	Models "../Model"

	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

//GetAllBooking Envia todos las reservar, formato->HTML
func GetAllBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []Models.Booking
	response, err := http.Get("http://localhost:8000/bookings")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &obj)

	tmpl, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Bookings/lstbookings.gohtml",
		"./View/templates/footer.gohtml",
	)

	wr.Header().Set("Content-Type", "text/html")

	err = tmpl.ExecuteTemplate(wr, "lstbookings", obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//GetAddEditBooking Envia un formulario con el objeto obtenido x el url, sino vacio
func GetAddEditBooking(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var booking Models.Booking
	//Si id tiene valor diferente a 0 entonces -> Edit, de lo contrario enviar 0;ver addeditbooking.gohtml href de Agregar
	if ps.ByName("id") != "0" {

		response, err := http.Get("http://localhost:8000/bookings/" + ps.ByName("id"))
		if err != nil {
			log.Fatal(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &booking)
	}
	tmpl, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Bookings/addeditbooking.gohtml",
		"./View/templates/footer.gohtml",
	)
	wr.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(wr, "addeditbooking", booking)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//PostPutBooking Inserta o actualiza una reserva
func PostPutBooking(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	var response *http.Request

	booking := &Models.Booking{
		UserID:   req.Form["BookingUserID"][0],
		FlightID: req.Form["FlightID"][0],
		PersonalSeat: Models.Seat{
			UserID: req.Form["SeatUserID"][0],
			Number: req.Form["Number"][0],
			Type:   req.Form["Type"][0],
		},
	}

	bookingBytes, err := json.Marshal(booking)
	if err != nil {
		fmt.Println(err)
		return
	}

	if req.Form["ID"][0] != "0" {
		response, _ = http.NewRequest("PUT", "http://localhost:8000/bookings/"+req.Form["ID"][0], bytes.NewReader(bookingBytes))
	} else {
		response, _ = http.NewRequest("POST", "http://localhost:8000/bookings", bytes.NewReader(bookingBytes))
	}
	new(http.Client).Do(response)

	GetAllBooking(wr, req, ps)
}

//DeleteBookingByID Elimina una reserva por ID, formato->JSON
func DeleteBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	response, _ := http.NewRequest("DELETE", "http://localhost:8000/bookings/"+ps.ByName("id"), nil)

	new(http.Client).Do(response)

	GetAllBooking(wr, req, ps)
}
