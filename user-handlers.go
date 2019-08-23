package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	users, err := userCollection.FindAll()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
	return
}

func GetUser(w http.ResponseWriter, req *http.Request, par httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(par.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := userCollection.FindOne(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
	return
}

func AddUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	decoder := json.NewDecoder(req.Body)

	var newUser User
	decodingError := decoder.Decode(&newUser)
	if decodingError != nil {
		http.Error(w, decodingError.Error(), http.StatusBadRequest)
		return
	}

	result, err := userCollection.InsertOne(newUser)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func UpdateUser(w http.ResponseWriter, req *http.Request, par httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(par.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	decoder := json.NewDecoder(req.Body)

	var newUser User
	decodingError := decoder.Decode(&newUser)
	if decodingError != nil {
		http.Error(w, decodingError.Error(), http.StatusBadRequest)
		return
	}

	res, err := userCollection.UpdateOne(id, newUser)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
	return
}

func DeleteUser(w http.ResponseWriter, req *http.Request, par httprouter.Params) {
	id, err := primitive.ObjectIDFromHex(par.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := userCollection.DeleteOne(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
	return
}
