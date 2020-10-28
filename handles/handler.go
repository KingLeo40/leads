package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/submission/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(GetSubmissions))).Methods(http.MethodGet)
	r.Handle("/submission", mw.Handler(http.HandlerFunc(CreateSubmission))).Methods(http.MethodPost)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "leads.submission.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
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
