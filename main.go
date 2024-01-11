// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharmakashish/go-crud-api/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", api.GetAllTodosHandler).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", api.GetTodoByIDHandler).Methods("GET")
	r.HandleFunc("/todos", api.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", api.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", api.DeleteTodoHandler).Methods("DELETE")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
