package users

import (
	"errors"
	"github.com/mnasruul/bookstore_utils-go/logger"
	"github.com/mnasruul/bookstore_utils-go/rest_errors"
	"github.com/mnasruul/go_rest_api_sqlserver/src/datasources/sqlserver/users_db"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT * FROM db_users.dbo.t_users;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() ([]User, rest_errors.RestErr) {
	rows, err := users_db.Client.Query(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return nil, rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer rows.Close()

	//rows:= stmt.Query()
	result := make([]User, 0)
	for rows.Next() {
		var user User
		if getErr := rows.Scan(&user.UserId, &user.UserName); getErr != nil {
			logger.Error("error when trying to get user by id", getErr)
			return nil, rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
		}
		result = append(result, user)
	}

	return result, nil
}
