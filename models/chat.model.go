package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rfauzi44/chat/db"
	"github.com/rfauzi44/chat/libs"
)

type Chat struct {
	ID          string    `json:"id"`
	SenderID    string    `json:"sender_id"`
	ReceiverID  string    `json:"receiver_id"`
	MessageText string    `json:"message_text"`
	CreatedAt   time.Time `json:"created_at"`
}

type ChatList struct {
	Sender      string    `json:"sender"`
	Receiver    string    `json:"receiver"`
	MessageText string    `json:"message_text"`
	CreatedAt   time.Time `json:"created_at"`
}

func AddChat(sender_id, receiver_id, message_text string) (libs.Response, error) {
	var res libs.Response

	uuid := uuid.New()
	id := uuid.String()

	conn := db.Connect()

	sqlStatement := "INSERT messages (id, sender_id, receiver_id, message_text, created_at) VALUES (?, ?, ?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, sender_id, receiver_id, message_text, time.Now())
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

func GetChatWith(sender_id, receiver_id string) (libs.Response, error) {
	var obj ChatList
	var objArr []ChatList
	var res libs.Response

	conn := db.Connect()

	sqlStatement := `SELECT u.username, u2.username, m.message_text, m.created_at
	FROM messages m
	JOIN users u ON m.sender_id = u.id
	JOIN users u2 ON m.receiver_id = u2.id
	WHERE (m.sender_id = ? AND m.receiver_id = ?) OR (m.sender_id = ? AND m.receiver_id = ?)
	ORDER BY m.created_at ASC;
	`

	rows, err := conn.Query(sqlStatement, sender_id, receiver_id, receiver_id, sender_id)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&obj.Sender,
			&obj.Receiver,
			&obj.MessageText,
			&obj.CreatedAt,
		)
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
