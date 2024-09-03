package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/5afar/Junior_Backend_To-Do/internal/database"
	"github.com/5afar/Junior_Backend_To-Do/internal/handlers"
)

func Run(){
	database.Init()

	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}