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

//HandleRequests - handles api requests
func HandleRequests() {
	personService.PopulateDb()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(constants.GetAllPerson, getAllPerson).Methods(constants.GET)
	myRouter.HandleFunc(constants.GetAllPersons, getAllPerson).Methods(constants.GET)
	myRouter.HandleFunc(constants.GetPersonById, getPersonById).Methods(constants.GET)
	myRouter.HandleFunc(constants.GetPersonByUserId, getPersonByUserId)
	myRouter.HandleFunc(constants.CreatePerson, createPerson).Methods(constants.POST)
	myRouter.HandleFunc(constants.UpdatePerson, updatePerson).Methods(constants.PUT)
	myRouter.HandleFunc(constants.DeletePerson, deletePerson).Methods(constants.DELETE)
	log.Fatal(http.ListenAndServe(constants.Host+":"+constants.PORT, myRouter))
}

//getAllPerson - fetch all person details
func getAllPerson(w http.ResponseWriter, r *http.Request) {
	logMsg.InfoLog("fetching all persons")
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

//getPersonById - fetch person details by id
func getPersonById(w http.ResponseWriter, r *http.Request) {
	logMsg.InfoLog("fetching person details by id")
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

//getPersonByUserId - fetch person details by userid
func getPersonByUserId(w http.ResponseWriter, r *http.Request) {
	logMsg.InfoLog("fetching person details by user id")
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

//createPerson - creates a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
	logMsg.InfoLog("creating a new person")
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !validation.IsValidPerson(p) {
		w.Write([]byte("person not valid"))
		return
	}
	personService.Add(p)
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.Write(reqBody)
}

//updatePerson - updates person details
func updatePerson(w http.ResponseWriter, r *http.Request) {
	logMsg.InfoLog("updating person details")
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

//deletePerson - delets a person
func deletePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	logMsg.InfoLog("deleting a person")

	_, err := personService.Delete(i)
	if err != nil {
		w.Write([]byte("failed to delete person"))
	}
	w.Write([]byte("successfully deleted a person"))

}
