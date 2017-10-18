package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	Models "../Model"

	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

//GetAllFlight Envia todos los vuelos, formato->HTML
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var obj []Models.Flight
	response, err := http.Get("http://localhost:8000/flights")

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
		"./View/Flights/lstflights.gohtml",
		"./View/templates/footer.gohtml",
	)
	wr.Header().Set("Content-Type", "text/html")

	err = tmpl.ExecuteTemplate(wr, "lstflights", obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//GetAddEditFlight Envia un formulario con el objeto obtenido x el url, sino vacio
func GetAddEditFlight(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var flight Models.Flight
	//Si id tiene valor diferente a 0 entonces -> Edit, de lo contrario enviar 0;ver addeditflight.gohtml href de Agregar
	if ps.ByName("id") != "0" {

		response, err := http.Get("http://localhost:8000/flights/" + ps.ByName("id"))
		if err != nil {
			log.Fatal(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &flight)
	}
	tmpl, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Flights/addeditflight.gohtml",
		"./View/templates/footer.gohtml",
	)
	wr.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(wr, "addeditflight", flight)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//PostPutFlight Inserta o actualiza un vuelo
func PostPutFlight(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	var response *http.Request

	parsePrice64, _ := strconv.ParseFloat(req.Form["Price"][0], 32)
	//chekar la diferencia entre pasarlo como memoria u objeto
	flight := &Models.Flight{
		AirplaneModel:  req.Form["AirplaneModel"][0],
		AirplaneNumber: req.Form["AirplaneNumber"][0],
		Price:          float32(parsePrice64),
		Depart: Models.Departure{
			Country: req.Form["DepartCountry"][0],
			City:    req.Form["DepartCity"][0],
			TD:      req.Form["DepartTD"][0],
			TA:      req.Form["DepartTA"][0],
		},
		Destin: Models.Destination{
			Country: req.Form["DestinCountry"][0],
			City:    req.Form["DestinCity"][0],
			TD:      req.Form["DestinTD"][0],
			TA:      req.Form["DestinTA"][0],
		},
		Seats: []Models.Seat{},
	}

	//Verificamos si existe un campo de Seats, si es asi lo rellenamos y lo agregamos al objeto flight.
	for i := 0; i < len(req.Form["UserID"]); i++ {
		seat := Models.Seat{
			UserID: req.Form["UserID"][i],
			Number: req.Form["Number"][i],
			Type:   req.Form["Type"][i],
		}
		flight.Seats = append(flight.Seats, seat)
	}

	flightBytes, err := json.Marshal(flight)
	if err != nil {
		fmt.Println(err)
		return
	}

	if req.Form["ID"][0] != "0" {
		response, _ = http.NewRequest("PUT", "http://localhost:8000/flights/"+req.Form["ID"][0], bytes.NewReader(flightBytes))
	} else {
		response, _ = http.NewRequest("POST", "http://localhost:8000/flights", bytes.NewReader(flightBytes))
	}
	new(http.Client).Do(response)

	GetAllFlight(wr, req, ps)
}

//DeleteFlightByID Elimina un vuelo por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	response, _ := http.NewRequest("DELETE", "http://localhost:8000/flights/"+ps.ByName("id"), nil)

	new(http.Client).Do(response)

	GetAllFlight(wr, req, ps)
}
