package controllers

import (
	"AnekaZoo/database"
	"AnekaZoo/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	database.Connector.Create(animal)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(animal)
}

//UpdateAnimalByID updates animal with respective ID
func UpdateAnimalByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	fmt.Print(r.Body)
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
