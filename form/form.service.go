package form

import (
	"github.com/google/uuid"
	"github.com/p-weisk/secret-form/config"
)

func CreateForm(content string, publicKey string) (Form, error) {
	form := Form{
		uuid.New(),
		content,
		publicKey,
	}
	err := form.persist(config.DB)
	return form, err
}

func GetForm(id uuid.UUID) (Form, error) {
	return retrieveForm(id, config.DB)
}
