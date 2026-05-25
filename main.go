package main

import (
	"log"
	"net/http"
	"os"

	"int-cont/handlers"
	"int-cont/store"

	"github.com/gorilla/mux"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	store.Init(redisHost, redisPort)

	r := mux.NewRouter()

	r.HandleFunc("/stack/{name}/push", handlers.Push).Methods("POST")
	r.HandleFunc("/stack/{name}/pop", handlers.Pop).Methods("POST")
	r.HandleFunc("/queue/{name}/enqueue", handlers.Enqueue).Methods("POST")
	r.HandleFunc("/queue/{name}/dequeue", handlers.Dequeue).Methods("POST")
	r.HandleFunc("/list/{name}/append", handlers.Append).Methods("POST")
	r.HandleFunc("/list/{name}/remove", handlers.Remove).Methods("POST")
	r.HandleFunc("/{type}/{name}", handlers.GetState).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
