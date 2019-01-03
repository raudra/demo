package controller


import (
		"net/http"
		"fmt"
    	"encoding/json"
    	"demo/app/models"
    	"log"
    	"io/ioutil"
)

type TodoParam struct {
	Name string `json:"site_key"`
	Id int   `json:"Id"`
}



type TodoParams struct{
	Tparams []TodoParam
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!....")
}

func Show(w http.ResponseWriter, r *http.Request){
	sites := model.AllSites()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)	
	todosJson, err := json.Marshal(sites)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	w.Write(todosJson)
}

func Create(w http.ResponseWriter, r*http.Request) {
	var todos []model.ToDo
	var params = new(TodoParams)
	
	body, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &params.Tparams)
	
	if err != nil {
		panic(err)
	}

	for _, params := range params.Tparams {
		todos = append(todos,model.New(params.Id, params.Name))
	}		

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	todosJson, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	w.Write(todosJson)
}



