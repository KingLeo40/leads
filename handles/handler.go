package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware("leads.submission.view", scrt, secureUrl, GetSubmissions)
	r.HandleFunc("/submission/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("leads.submission.create", scrt, secureUrl, CreateSubmission)
	r.HandleFunc("/submission", create).Methods(http.MethodPost)

	//leads.submission.search

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "leads.submission.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
