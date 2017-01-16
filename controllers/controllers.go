package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	m "github.com/noliva/api-maps/models"
	r "github.com/noliva/api-maps/models/repositories"
	"time"
)

func CreateMap(w http.ResponseWriter, req *http.Request) {
	var newMyMap m.MyMap
	now := time.Now()
	newMyMap.CreatedAt = now.Format("2006-01-02 15:04:05")
	_ = json.NewDecoder(req.Body).Decode(&newMyMap)
	r.Create(newMyMap)

	json.NewEncoder(w).Encode(&newMyMap)
}

func GetMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	myMap := r.FindMapById(params["id"])

	json.NewEncoder(w).Encode(&myMap)
}

func GetMaps(w http.ResponseWriter, req *http.Request) {
	myMap := r.FindAllMaps()

	json.NewEncoder(w).Encode(&myMap)
}

func DeleteMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	myMap := r.FindMapById(params["id"])

	r.Delete(myMap)
}