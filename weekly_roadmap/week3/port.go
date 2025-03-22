package week3

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Week3_Port() {
	service := NewStudentService()
	controller := NewStudentController(service)

	r := mux.NewRouter()
	r.HandleFunc("/students", controller.GetAll).Methods("GET")
	r.HandleFunc("/students/{id}", controller.GetByID).Methods("GET")
	r.HandleFunc("/students", controller.Add).Methods("POST")
	r.HandleFunc("/students/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/students/{id}", controller.Delete).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
