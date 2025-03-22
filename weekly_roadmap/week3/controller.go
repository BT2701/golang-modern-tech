package week3

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentController struct {
	service *StudentService
}

func NewStudentController(service *StudentService) *StudentController {
	return &StudentController{service: service}
}

func (c *StudentController) GetAll(w http.ResponseWriter, r *http.Request) {
	students := c.service.GetAll()
	json.NewEncoder(w).Encode(students)
}

func (c *StudentController) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := c.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func (c *StudentController) Add(w http.ResponseWriter, r *http.Request) {
	var student Student
	json.NewDecoder(r.Body).Decode(&student)

	student = c.service.Add(student)
	json.NewEncoder(w).Encode(student)
}

func (c *StudentController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student Student
	json.NewDecoder(r.Body).Decode(&student)

	student, err = c.service.Update(id, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func (c *StudentController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
