package system

import (
	"encoding/json"
	"log"
	"net/http"
	// "fmt"

	"github.com/gorilla/mux"

	"server/controllers"
	"server/model"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := "pong"
	json.NewEncoder(w).Encode(response)
	log.Print(response)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := controllers.Read()
	json.NewEncoder(w).Encode(response)
}

func ReadOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	response := controllers.ReadOne(params["id"])
	json.NewEncoder(w).Encode(response)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	controllers.Delete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func CreateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task model.Student
	_ = json.NewDecoder(r.Body).Decode(&task)
	controllers.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task model.Student
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&task)
	controllers.Update(&task, params["id"])
	json.NewEncoder(w).Encode(task)
}
