package form

import (
	"database/sql"
	"github.com/google/uuid"
)

const createFormQuery = "INSERT INTO secretform.forms (ID, Content, PublicKey, Open) VALUES (?,?,?,?);"

const getFormQuery = "SELECT ID, Content, PublicKey, Open FROM secretform.forms;"

const closeFormQuery = "UPDATE secretform.forms SET Open = false WHERE ID = ?;"

type Form struct {
	ID			uuid.UUID
	Content		string
	PublicKey	string
	Open		bool
}

// Persist a form to the given database
func (f Form) persist(db *sql.DB) error {
	id := f.ID.String()
	_, err := db.Exec(createFormQuery, id, f.Content, f.PublicKey, f.Open)
	return err
}

// Get a form
func retrieveForm(fid string, db *sql.DB) (Form, error) {
	result := Form{}
	rows, rerr := db.Query(getFormQuery, fid)
	if rerr != nil {
		return result, rerr
	}
	for rows.Next() {
		var id string
		err := rows.Scan(&id, &result.Content, &result.PublicKey, &result.Open)
		if err != nil {
			return result, err
		}
		result.ID = uuid.MustParse(id)
	}

	return result, nil
}

// Close a form
func closeForm(id string, db *sql.DB) error {
	_, err := db.Exec(closeFormQuery, id)
	return err
}