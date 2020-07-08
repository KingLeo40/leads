package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Submissions husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{Submissions: husk.NewTable(Submission{})}
}

func Shutdown() {
	ctx.Submissions.Save()
}

func GetSubmissions(page, size int) (husk.Collection, error) {
	return ctx.Submissions.Find(page, size, husk.Everything())
}

func CreateSubmission(obj Submission) error {
	_, err := ctx.Submissions.Create(obj)

	if err != nil {
		return err
	}

	return ctx.Submissions.Save()
}
