package users

import (
	"github.com/mnasruul/bookstore_oauth-go/rest_errors"
	"strings"
)


const (
	StatusActive = "active"
)

type User struct {
	UserId   int64  `json:"userid"`
	UserName string `json:"username"`
}

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.UserName = strings.TrimSpace(user.UserName)
	//user.LastName = strings.TrimSpace(user.LastName)
	//
	//user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//if user.Email == "" {
	//	return rest_errors.NewBadRequestError("invalid email address")
	//}
	//
	//user.Password = strings.TrimSpace(user.Password)
	//if user.Password == "" {
	//	return rest_errors.NewBadRequestError("invalid password")
	//}
	return nil
}