package form

import (
	"github.com/google/uuid"
	"github.com/p-weisk/secret-form/config"
)

func CreateForm(content string, publicKey string) error {
	form := Form{
		uuid.New(),
		content,
		publicKey,
		true,
	}
	err := form.persist(config.DB)
	return err
}

func GetForm(id string) (Form, error) {
	return retrieveForm(id, config.DB)
}

func CloseForm(id string) error {
	return closeForm(id, config.DB)
}
