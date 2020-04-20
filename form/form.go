package form

import (
	"database/sql"
	"github.com/google/uuid"
)

const createFormQuery = "INSERT INTO secretform.forms (ID, Content, PublicKey) VALUES (?,?,?);"

const getFormQuery = "SELECT ID, Content, PublicKey FROM secretform.forms;"

type Form struct {
	ID			uuid.UUID
	Content		string
	PublicKey	string
}

// Persist a form to the given database
func (f Form) persist(db *sql.DB) error {
	id := f.ID.String()
	_, err := db.Exec(createFormQuery, id, f.Content, f.PublicKey)
	return err
}

// Get a form
func retrieveForm(fid uuid.UUID, db *sql.DB) (Form, error) {
	result := Form{}
	rows, rerr := db.Query(getFormQuery, fid.String())
	if rerr != nil {
		return result, rerr
	}
	for rows.Next() {
		var id string
		err := rows.Scan(&id, &result.Content, &result.PublicKey)
		if err != nil {
			return result, err
		}
		result.ID = uuid.MustParse(id)
	}

	return result, nil
}
