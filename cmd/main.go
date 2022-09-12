package main

import (
	"fmt"
	"log"
	"net/http"

	//"math/rand"

	"go-goland-api/main.go/server"
)

var granId = 2

//var tasks = []domain.Task{
//	{
//		ID:      1,
//		Nombre:  "task one",
//		Content: "some content",
//	},
//	{
//		ID:      2,
//		Nombre:  "task two",
//		Content: "more content",
//	},
//}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")

}

func main() {
	const port string = ":8888"
	httpServer := server.NewHTTPRouter(port)
	router := httpServer.SetupRouter()

	log.Println("Server listening on port", port)
	if err := httpServer.Run(router); err != nil {
		log.Fatalln(err)
	}

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", indexRoute)
	// router.HandleFunc("/tasks", getTasks).Methods("GET")
	// router.HandleFunc("/tasks/{ID}", getTask).Methods("GET")
	// router.HandleFunc("/tasks", createTask).Methods("POST")
	// router.HandleFunc("/tasks/{ID}", putTask).Methods("PUT")
	// router.HandleFunc("/tasks/{ID}", deleteTask).Methods("DELETE")
	// log.Fatal(http.ListenAndServe(":3000", router))

}

// GET ALL TASKS
// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// } // GET ONE TASK
// func getTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	aux := mux.Vars(r)
// 	auxID, err := strconv.Atoi(aux["ID"])

// 	if err != nil {
// 		fmt.Fprintf(w, "Id invalido")
// 		return
// 	}

// 	for _, task := range tasks {
// 		if task.ID == auxID {
// 			json.NewEncoder(w).Encode(task)
// 			return
// 		}

// 	}
// 	json.NewEncoder(w).Encode(domain.ResponseInfo{
// 		Status: http.StatusBadRequest,
// 		Data:   "error ",
// 	})

// }

// // CREATE A NEW TASK
// func createTask(w http.ResponseWriter, r *http.Request) {
// 	var newTask domain.Task
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalido")
// 	}
// 	json.Unmarshal(reqBody, &newTask)
// 	granId = granId + 1
// 	newTask.ID = granId
// 	tasks = append(tasks, newTask)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newTask)

// }
// func putTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	aux := mux.Vars(r)
// 	auxID, err := strconv.Atoi(aux["ID"])
// 	var updatedTask domain.Task
// 	if err != nil {
// 		fmt.Fprintf(w, "Id invalido")
// 		return
// 	}
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Favor de ingresar datos validos")

// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, task := range tasks {
// 		if task.ID == auxID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			updatedTask.ID = task.ID
// 			tasks = append(tasks, updatedTask)
// 			fmt.Fprintf(w, "update complete")
// 		}
// 	}

// }

// // DELETE TASKS
// func deleteTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	aux := mux.Vars(r)
// 	auxID, err := strconv.Atoi(aux["ID"])

// 	if err != nil {
// 		fmt.Fprintf(w, "Id invalido")
// 		return
// 	}

// 	for index, task := range tasks {
// 		if task.ID == auxID {
// 			tasks = append(tasks[:index], tasks[index+1:]...)
// 			fmt.Fprintf(w, "Tasks deleted")
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(domain.ResponseInfo{
// 		Status: http.StatusBadRequest,
// 		Data:   "Id no encontrado ",
// 	})

// }
