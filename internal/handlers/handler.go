package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"person/internal/constants"
	"person/internal/database"
	logMsg "person/internal/log"
	"person/internal/model"
	"person/internal/services"
	"person/internal/validation"
	"strconv"

	"github.com/gorilla/mux"
)

var conn, _ = database.ConnectToDB(database.EmptyDbConstruct{})
var personService = services.NewPersonService(conn)

func HandleRequests() {
	personService.PopulateDb()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getAllPerson).Methods(constants.GET)
	myRouter.HandleFunc("/persons", getAllPerson).Methods(constants.GET)
	myRouter.HandleFunc("/persons/{id}", getPersonById).Methods(constants.GET)
	myRouter.HandleFunc("/persons/user/{id}", getPersonByUserId)
	myRouter.HandleFunc("/person", createPerson).Methods(constants.POST)
	myRouter.HandleFunc("/persons", updatePerson).Methods(constants.PUT)
	myRouter.HandleFunc("/persons/{id}", deletePerson).Methods(constants.DELETE)
	log.Fatal(http.ListenAndServe(constants.Host+":"+constants.PORT, myRouter))
}
func getAllPerson(w http.ResponseWriter, r *http.Request) {
	logMsg.WriteInfoLog(" reading persons")
	persons, err := personService.GetPersons()
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	if err != nil {
		panic(err.Error())
	}
	newFsConfigBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.Write(newFsConfigBytes)

}
func getPersonById(w http.ResponseWriter, r *http.Request) {
	logMsg.WriteInfoLog(" reading person by id")
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	persons, err := personService.GetPersonById(i)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	if err != nil {
		panic(err.Error())
	}
	newFsConfigBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.Write(newFsConfigBytes)
}

func getPersonByUserId(w http.ResponseWriter, r *http.Request) {
	logMsg.WriteInfoLog(" reading person by user id")
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	persons, err := personService.GetPersonByUserId(i)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	if err != nil {
		panic(err.Error())
	}
	newFsConfigBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	w.Write(newFsConfigBytes)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	logMsg.WriteInfoLog(" creating the person")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("person ->", p)
	if !validation.IsValidPerson(p) {
		w.Write([]byte("person not valid"))
		return
	}
	personService.Add(p)
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.Write(reqBody)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	logMsg.WriteInfoLog(" updating the person")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	personService.Update(p)
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.Write(reqBody)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	logMsg.WriteInfoLog(" Deleting the person")

	_, err := personService.Delete(i)
	if err != nil {
		w.Write([]byte("Failure to delete"))
	}
	w.Write([]byte("successfully deleted"))

}
