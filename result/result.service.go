package result

import (
	"github.com/google/uuid"
	"github.com/p-weisk/secret-form/answer"
	"github.com/p-weisk/secret-form/form"
)

func GetResult(id uuid.UUID) (Result, error) {

	form, err1 := form.GetForm(id)
	if err1 != nil {
		return Result{}, err1
	}

	answers, err2 := answer.GetAnswers(id)
	if err2 != nil {
		return Result{}, err2
	}

	return Result{
		form,
		answers,
	},
	nil
}
