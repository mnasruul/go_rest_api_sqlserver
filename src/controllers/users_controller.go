package controllers

import (
	"github.com/gorilla/mux"
	"github.com/mnasruul/bookstore_items-api/src/utils/http_utils"
	"github.com/mnasruul/bookstore_oauth-go/rest_errors"
	"github.com/mnasruul/go_rest_api_sqlserver/src/services"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	//Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	//Search(w http.ResponseWriter, r *http.Request)
}

type usersController struct {}

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}
func (cont *usersController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["user_id"])
	userId, idErr := getUserId(strings.TrimSpace(vars["user_id"]))
	log.Println(userId)
	if idErr != nil {
		http_utils.RespondError(w, idErr)
		return
	}
	user, err := services.UsersService.GetUser(userId)
	log.Println(user)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, user)
}