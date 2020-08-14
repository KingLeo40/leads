package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/leads/core"
	"log"
	"net/http"
)

func GetSubmissions(w http.ResponseWriter, r *http.Request) {
	p, s := drx.GetPageData(r)

	res, err := core.GetSubmissions(p, s)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(res))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	obj := core.Submission{}
	err := drx.JSONBody(r, &obj)

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

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
