package models

import (
	"golang-crud-test/db"
	"net/http"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id		int	`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	FullName	string	`json:"full_name"`
	Photo		string	`json:"photo"`
}

func FetchAllUser() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()

	fmt.Println(con)


	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Password, &obj.FullName, &obj.Photo)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreUser(id int, username string, rawPassword string, fullname string) (Response, error) {
	var res Response

	con := db.CreateCon()

	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)

	sqlStatement := "INSERT INTO users (username, password, full_name) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, password, fullname)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func UpdateUser(id int, username string, rawPassword string, fullname string) (Response, error) {
	var res Response

	con := db.CreateCon()

	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)

	sqlStatement := "UPDATE users SET username = ?, password = ?, full_name = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, password, fullname, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
