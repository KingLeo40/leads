package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/leads/core"
	"log"
	"net/http"
)

func GetSubmissions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	p, s := ctx.GetPageData()

	res, err := core.GetSubmissions(p, s)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(res))
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	obj := core.Submission{}
	err := ctx.Body(&obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.CreateSubmission(obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON("Saved"))
}
