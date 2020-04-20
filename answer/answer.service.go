package answer

import (
	"github.com/google/uuid"
	"github.com/p-weisk/secret-form/config"
)

func CreateAnswer(data string, form uuid.UUID) error {
	answer := Answer{
		uuid.New(),
		data,
		form,
	}
	err := answer.persist(config.DB)
	return err
}

func GetAnswers(fid uuid.UUID) ([]Answer, error) {
	return retrieveAnswers(fid, config.DB)
}