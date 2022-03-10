package controllers

import (
	"AnekaZoo/database"
	"AnekaZoo/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

// Result is an array of animal
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//GetAllAnimal get all animal data
func GetAllAnimal(w http.ResponseWriter, r *http.Request) {
	var animals []entity.Animal
	database.Connector.Find(&animals) //tabel name animals
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(animals)
}

//GetAnimalByID returns animal with specific ID
func GetAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var animal entity.Animal
	database.Connector.First(&animal, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

/*
Create Animal Func
*/
func CreateAnimal(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var animal entity.Animal
	json.Unmarshal(requestBody, &animal)

	if duplicateError := database.Connector.Select("Name", "Class", "Legs").Create(&animal).Error; duplicateError != nil {

		errorResponse := Result{Code: 500, Data: animal, Message: "Duplicate name animal"}
		result, err := json.Marshal(errorResponse)
		if err != nil {
			http.Error(w, duplicateError.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
		return
	}

	res := Result{Code: 200, Data: animal, Message: "Animal has been created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(result)
	w.Write(result)
}

//UpdateAnimalByID updates animal with respective ID
func UpdateAnimalByID(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		err.Error()
		return

	}
	var animal entity.Animal
	json.Unmarshal(requestBody, &animal)
	database.Connector.Save(&animal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(animal)
}

//DeletAnimalByID delete's animal with specific ID
func DeletAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var animal entity.Animal
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&animal)
	w.WriteHeader(http.StatusNoContent)
}
