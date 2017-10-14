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

//GetAllUser Envia todos los usuarios, formato->HTML
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var obj []Models.User
	response, err := http.Get("http://localhost:8000/users")

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
		"./View/Users/lstusers.gohtml",
		"./View/templates/footer.gohtml",
	)

	wr.Header().Set("Content-Type", "text/html")

	err = tmpl.ExecuteTemplate(wr, "lstusers", obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//GetAddEditUser Envia un formulario conel objeto obtenido x el url, sino vacio
func GetAddEditUser(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var user Models.User
	//Si id tiene valor diferente a 0 entonces -> Edit, de lo contrario enviar 0;ver addedituser.gohtml href de Agregar
	if ps.ByName("id") != "0" {

		response, err := http.Get("http://localhost:8000/users/" + ps.ByName("id"))
		if err != nil {
			log.Fatal(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &user)
	}
	tmpl, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Users/addedituser.gohtml",
		"./View/templates/footer.gohtml",
	)
	wr.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(wr, "addedituser", user)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

//PostPutUser Inserta o actualiza un usuario
func PostPutUser(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	var response *http.Request

	parseTotal64, _ := strconv.ParseFloat(req.Form["Total"][0], 32)

	user := &Models.User{
		FirstName:      req.Form["FirstName"][0],
		LastName:       req.Form["LastName"][0],
		PassportType:   req.Form["PassportType"][0],
		PassportNumber: req.Form["PassportNumber"][0],
		Email:          req.Form["Email"][0],
		Password:       req.Form["Password"][0],
		PersonalCard: Models.Payment{
			Card:           req.Form["Card"][0],
			CardNumber:     req.Form["CardNumber"][0],
			CSC:            req.Form["CSC"][0],
			Total:          float32(parseTotal64),
			ExpirationDate: req.Form["ExpirationDate"][0],
		},
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	// var user Models.User

	// user.FirstName=

	if req.Form["ID"][0] != "0" {
		log.Println("Editar")
	} else {
		response, _ = http.NewRequest("POST", "http://localhost:8000/users", bytes.NewReader(userBytes))
	}
	new(http.Client).Do(response)

	GetAllUser(wr, req, ps)
}

// //PutUserByID Actualiza un Documento User
// func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

// }

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	response, _ := http.NewRequest("DELETE", "http://localhost:8000/users/"+ps.ByName("id"), nil)

	new(http.Client).Do(response)

	GetAllUser(wr, req, ps)
}

//RenderIndex No es de User, solo se utiliza para mostrar el menu
func RenderIndex(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	wr.Header().Set("Content-Type", "text/html")

	wr.Write([]byte(
		`
		<h3>Menu</h3>
		<ul>
			<li><a href="http://localhost:8001/lstusers.html">Users</a></li>
			<li><a href="http://localhost:8001/lstflights.html">Flights</a></li>
			<li><a href="http://localhost:8001/lstbookings.html">Bookings</a></li>
	  	</ul> 
		`),
	)
}
