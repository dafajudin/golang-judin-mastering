package models

import (
	"database/sql"
	"fmt"
	"go-hero/db"
	"go-hero/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	//searching data for username
	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	//compare password
	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password Doesn't Match")
		return false, err
	}
	return true, nil
}
