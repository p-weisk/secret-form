package result

import (
	"github.com/p-weisk/secret-form/answer"
	"github.com/p-weisk/secret-form/form"
)

type Result struct {
	Form	form.Form
	Answers	[]answer.Answer
}
