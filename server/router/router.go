package router

import (
	"github.com/gorilla/mux"

	middleware "server/system/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/ping", middleware.Ping).Methods("GET", "OPTIONS")
	router.HandleFunc("/read", middleware.ReadAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/readOne/{id}", middleware.ReadOneStudent).Methods("GET", "OPTIONS")
	router.HandleFunc("/delete/{id}", middleware.DeleteOne).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/create", middleware.CreateOne).Methods("POST", "OPTIONS")
	router.HandleFunc("/update/{id}", middleware.UpdateOne).Methods("PUT", "OPTIONS")
	return router
}
