package controller

import (
		"net/http"
    	"encoding/json"
    	"demo/app/models"
    	"log"
)

func Sites(w http.ResponseWriter, r *http.Request){
	sites := model.AllSites()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)	
	todosJson, err := json.Marshal(sites)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	w.Write(todosJson)
}
func ShowSite(w http.ResponseWriter, r *http.Request) {
	
}
