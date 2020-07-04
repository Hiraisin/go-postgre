package models

import (
	"database/sql"
	"fmt"

	"github.com/hiraisin/go-postgre/config"
)

func CheckLogin(username, password string) error {
	var obj User
	var pwd string

	con := config.CreateCon()

	sqlStatement := "SELECT id,username,password FROM users where username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return err
	}

	if err != nil {
		fmt.Println("Query Error")
		return err
	}

	match := config.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match")
		fmt.Println(pwd, password)
		return err
	}

	return nil
}
