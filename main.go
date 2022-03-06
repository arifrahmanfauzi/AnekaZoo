package main

import (
	"AnekaZoo/controllers"
	"AnekaZoo/database"
	"AnekaZoo/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8000")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateAnimal).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllAnimal).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetAnimalByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateAnimalByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletAnimalByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "anekazoo",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Animal{})
}
