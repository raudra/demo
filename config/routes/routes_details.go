package routes

import (
			"net/http"
			"demo/app/controllers"
		)

type Route struct{
	Name 		string
	Method		string
	Pattern		string
	HandlerFunc     http.HandlerFunc
}


type  Routes []Route

var routes = Routes{
	Route{
		"Controller: SiteController Action: sites", 
		"GET",  
		"/sites", 
		controller.Sites,
	},
	Route{
		"Controller: TodoController Action: Show",
		"GET",
		"/todos/{id}",
		controller.Show,
	},
	Route{
		"Create",
		"POST",
		"/todos",
		controller.Create,
	},
}
