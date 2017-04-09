package main

import (
	"net/http"

	"github.com/markcheno/go-vue-starter/controllers"
	"github.com/markcheno/go-vue-starter/models"
	"github.com/rs/cors"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/negroni"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return controllers.SigningKey, nil
	},
	UserProperty:  "user",
	Debug:         false,
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {

	n := negroni.Classic()

	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.LogMode(true)

	users, _ := models.NewUserState(db)
	controller := controllers.NewController(db, users)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost"},
	})

	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("./webpack/dist/"))).Methods("GET")
	mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./webpack/dist/static/js/"))))
	mux.HandleFunc("/signup", controller.Signup).Methods("POST")
	mux.HandleFunc("/login", controller.Login).Methods("POST")
	mux.HandleFunc("/api/random-quote", controller.Quote).Methods("GET")
	mux.Handle("/api/protected/random-quote", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(controller.SecretQuote)),
	))
	mux.Handle("/test-protected", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(controller.TestProtected)),
	))

	n.Use(c)
	n.UseHandler(mux)

	n.Run(":3000")
}
