package models

import (
	"database/sql"
	"fmt"

	"github.com/hiraisin/go-postgre/config"
)

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := config.CreateCon()

	sqlStatement := "SELECT id,username,password FROM users where username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match := config.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match")
		fmt.Println(pwd, password)
		return false, err
	}

	return true, nil
}
