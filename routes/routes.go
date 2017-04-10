package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markcheno/go-vue-starter/api"
	"github.com/markcheno/go-vue-starter/auth"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	// quotes
	q := a.PathPrefix("/quote").Subrouter()
	q.HandleFunc("/random", api.Quote).Methods("GET")
	q.Handle("/protected/random", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	))

	return mux
}
