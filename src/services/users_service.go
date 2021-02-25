package services

import (
	"github.com/mnasruul/bookstore_oauth-go/rest_errors"
	"github.com/mnasruul/go_rest_api_sqlserver/src/domain/users"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser() (users.Users, rest_errors.RestErr)
	//CreateUser(users.User) (*users.User, rest_errors.RestErr)
	//UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	//DeleteUser(int64) rest_errors.RestErr
	//SearchUser(string) (users.Users, rest_errors.RestErr)
	//LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}
type usersService struct{}

func (s *usersService) GetUser() (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.Get()
}