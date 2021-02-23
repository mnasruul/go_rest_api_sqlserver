package app

import (
	"github.com/mnasruul/go_rest_api_sqlserver/src/controllers"
	"net/http"
)

func mapUrls(){
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}",controllers.UsersController.Get).Methods(http.MethodGet)
	router.HandleFunc("/users",controllers.UsersController.Get ).Methods(http.MethodPost)
}