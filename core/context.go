package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Submissions husk.Tabler
}

var ctx context

func NewContext() {
	ctx = context{Submissions: husk.NewTable(Submission{}, serials.GobSerial{})}
	return NewAPI(ctx, 10)
}

func Shutdown() {
	ctx.Submissions.Save()
}
