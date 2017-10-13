package controller

//Restful - User
import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

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

	testTemplate, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Users/lstusers.gohtml",
		"./View/templates/footer.gohtml",
	)

	wr.Header().Set("Content-Type", "text/html")

	err = testTemplate.ExecuteTemplate(wr, "lstusers", obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

// //PostUser Inserta un nuevo vuelo
// func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

// }

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
