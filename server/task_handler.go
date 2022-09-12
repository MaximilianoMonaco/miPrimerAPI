package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"go-goland-api/main.go/domain"
	"go-goland-api/main.go/repositories"

	"github.com/gorilla/mux"
)

var tasks []domain.Task

type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

var granId = 2

type Handler interface {
	getTask(w http.ResponseWriter, r *http.Request)
	getTasks(w http.ResponseWriter, r *http.Request)
	updatedTask(w http.ResponseWriter, r *http.Request)
	putTask(w http.ResponseWriter, r *http.Request)
	deleteTask(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo repositories.TaskRepository
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
} // GET ONE TASK
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	aux := mux.Vars(r)
	auxID, err := strconv.Atoi(aux["ID"])

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}

	for _, task := range tasks {
		if task.ID == auxID {
			json.NewEncoder(w).Encode(task)
			return
		}

	}
	json.NewEncoder(w).Encode(domain.ResponseInfo{
		Status: http.StatusBadRequest,
		Data:   "error ",
	})

}

// CREATE A NEW TASK
func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var newTask domain.Task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalido")
	}
	json.Unmarshal(reqBody, &newTask)
	granId = granId + 1
	newTask.ID = granId
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}
func putTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	aux := mux.Vars(r)
	auxID, err := strconv.Atoi(aux["ID"])
	var updatedTask domain.Task
	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Favor de ingresar datos validos")

	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == auxID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = task.ID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "update complete")
		}
	}

}

// DELETE TASKS
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	aux := mux.Vars(r)
	auxID, err := strconv.Atoi(aux["ID"])

	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}

	for index, task := range tasks {
		if task.ID == auxID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(w, "Tasks deleted")
			return
		}
	}
	json.NewEncoder(w).Encode(domain.ResponseInfo{
		Status: http.StatusBadRequest,
		Data:   "Id no encontrado ",
	})

}
