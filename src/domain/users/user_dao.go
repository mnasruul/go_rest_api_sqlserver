package users

import (
	"errors"
	"github.com/mnasruul/bookstore_oauth-go/rest_errors"
	"github.com/mnasruul/go_rest_api_sqlserver/src/datasources/sqlserver/db"
	"log"
)

const (
	//queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT * FROM t_users where user_id=?;"
	//queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	//queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	//queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	//queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() rest_errors.RestErr {
	stmt, err := sqlserver.Client.Prepare(queryGetUser)
	log.Println("DAO err:", err)
	if err != nil {
		//logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	log.Println("DAO stmt:", stmt)
	defer stmt.Close()
	log.Println(user)
	result := stmt.QueryRow(user.UserId)

	if getErr := result.Scan(&user.UserId, &user.UserName); getErr != nil {
		//logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil
}