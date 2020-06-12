package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Submissions husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{Submissions: husk.NewTable(Submission{}, serials.GobSerial{})}
}

func Shutdown() {
	ctx.Submissions.Save()
}

func GetSubmissions(page, size int) (husk.Collection, error) {
	return ctx.Submissions.Find(page, size, husk.Everything())
}

func CreateSubmission(obj Submission) error {
	cs := ctx.Submissions.Create(obj)

	if cs.Error != nil {
		return cs.Error
	}

	return ctx.Submissions.Save()
}
