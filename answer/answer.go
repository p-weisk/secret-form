package answer

import (
	"database/sql"
	"github.com/google/uuid"
)

const createAnswerQuery = "INSERT INTO secretform.answers (ID, Data, Form) VALUES (?,?,?);"

const getAnswersQuery = "SELECT ID, Data, Form, FROM secretform.answers WHERE Form = ?;"

type Answer struct {
	ID		uuid.UUID
	Data	string
	Form	uuid.UUID
}

func (a Answer) persist(db *sql.DB) error {
	id := a.ID.String()
	_, err := db.Exec(createAnswerQuery, id, a.Data, a.Form.String())
	return err
}

func retrieveAnswers(fid uuid.UUID, db *sql.DB) (a []Answer, err error) {
	a = []Answer{}
	rows, rerr := db.Query(getAnswersQuery, fid.String())
	if rerr != nil {
		return nil, rerr
	}
	for rows.Next() {
		var id string
		var f string
		an := Answer{}
		err := rows.Scan(&id, &an.Data, &f)
		if err != nil {
			return a, err
		}
		an.ID = uuid.MustParse(id)
		an.Form = uuid.MustParse(f)
	}

	return a, nil
}
