package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/markcheno/go-vue-starter/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

// GetToken create a JWT
func GetToken(user *models.User) string {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signedToken, _ := token.SignedString(SigningKey)
	return signedToken
}

// Signup -
func (c *Controller) Signup(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	formdata := UserForm{}
	err := decoder.Decode(&formdata)

	if err != nil || formdata.Username == "" || formdata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if c.users.HasUser(formdata.Username) {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	user := c.users.AddUser(formdata.Username, formdata.Password)

	js, err := json.Marshal(&Token{GetToken(user)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UserForm -
type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login -
func (c *Controller) Login(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	formdata := UserForm{}
	err := decoder.Decode(&formdata)

	if err != nil || formdata.Username == "" || formdata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	//var user = _.find(users, userScheme.userSearch);
	user := c.users.FindUser(formdata.Username)
	if user.Username == "" {
		http.Error(w, "username not found", http.StatusBadRequest)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formdata.Password)) != nil {
		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(&Token{GetToken(user)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// GetUser -
func (c *Controller) GetUser(req *http.Request) *models.User {
	u := context.Get(req, "user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := c.users.FindUserByUUID(u["uuid"].(string))
	return user
}
