package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type context struct {
	Submissions husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{Submissions: husk.NewTable(Submission{})}
}

func Shutdown() {
	ctx.Submissions.Save()
}

func GetSubmissions(page, size int) (records.Page, error) {
	return ctx.Submissions.Find(page, size, op.Everything())
}

func CreateSubmission(obj Submission) error {
	_, err := ctx.Submissions.Create(obj)

	if err != nil {
		return err
	}

	return ctx.Submissions.Save()
}
