package server

import (
	"encoding/json"
	"net/http"

	"go-goland-api/main.go/repositories"

	"github.com/gorilla/mux"
)

type HTTPRouter interface {
	SetupRouter() *mux.Router
	Run(router http.Handler) error
}

type httpRouter struct {
	port string
}

func (r *httpRouter) SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods("GET")

	dbConn, err := repositories.GetConnectionDB()
	if err != nil {
		panic("error db")
	}

	taskRepository := repositories.NewTaskRepository(dbConn)

	taskHandler := newHandler(taskRepository)

	router.HandleFunc("/tasks", taskHandler.getTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.updatedTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.putTask).Methods("PUT")
	router.HandleFunc("/ttasks/{id}", taskHandler.deleteTask).Methods("DELETE")

	return router
}
func (r *httpRouter) Run(router http.Handler) error {
	return http.ListenAndServe(r.port, router)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   "pong",
	})
}
