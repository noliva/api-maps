package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"

	m "github.com/noliva/api-maps/models"
	r "github.com/noliva/api-maps/models/repositories"
)

func CreateMap(w http.ResponseWriter, req *http.Request) {
	var newAreaMap m.MyMap
	_ = json.NewDecoder(req.Body).Decode(&newAreaMap)
	newAreaMap.ID = uuid.NewV4().String()
	r.Insert(newAreaMap)
	json.NewEncoder(w).Encode(newAreaMap)
}

func GetMap(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	newAreaMap := r.FindMapById(params["id"])

	json.NewEncoder(w).Encode(&newAreaMap)
}

func GetMaps(w http.ResponseWriter, req *http.Request) {
	areaMaps := r.FindAllMaps()

	json.NewEncoder(w).Encode(&areaMaps)
}