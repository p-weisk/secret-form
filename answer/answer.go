package answer

import (
	"database/sql"
	"github.com/google/uuid"
)

const createAnswerQuery = "INSERT INTO secretform.answers (ID, Data, Form) VALUES (?,?,?);"

const getAnswersQuery = "SELECT ID, Data, Form, FROM secretform.answers WHERE Form = ?;"

type Answer struct {
	ID			uuid.UUID
	Data		string
	Form		uuid.UUID
}

func (a Answer) persist(db *sql.DB) error {
	id := a.ID.String()
	_, err := db.Exec(createAnswerQuery, id, a.Data, a.Form.String())
	return err
}
