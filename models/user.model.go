package models

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/rfauzi44/chat-api/db"
	"github.com/rfauzi44/chat-api/libs"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func GetAllUser() (libs.Response, error) {
	var obj User
	var objArr []User
	var res libs.Response

	conn := db.Connect()

	sqlStatement := "SELECT * FROM users"

	rows, err := conn.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Username)
		if err != nil {
			return res, err
		}

		objArr = append(objArr, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = objArr
	return res, nil

}

func AddUser(username string) (libs.Response, error) {
	var res libs.Response

	uuid := uuid.New()
	id := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT users (id, username) VALUES (?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, username)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"last_inserted_id": id,
	}

	return res, nil
}
