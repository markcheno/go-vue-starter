package controllers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
	"github.com/markcheno/go-vue-starter/models"
)

// Controller -
type Controller struct {
	db    *gorm.DB
	users *models.UserState
}

// NewController -
func NewController(db *gorm.DB, userstate *models.UserState) *Controller {
	return &Controller{db: db, users: userstate}
}

// Token -
type Token struct {
	IDToken string `json:"id_token"`
}

// SigningKey Set up a global string for our secret
var SigningKey = []byte("knrjkevdckjh")

// TestProtected -
func (c *Controller) TestProtected(w http.ResponseWriter, req *http.Request) {

	user := c.GetUser(req)
	u := context.Get(req, "user").(*jwt.Token).Claims.(jwt.MapClaims)

	fmt.Fprintf(w, "This is an authenticated request\n")
	fmt.Fprintf(w, "Claim content:\n")
	fmt.Fprintf(w, "user.username: %s\n", user.Username)
	for k, v := range u {
		fmt.Fprintf(w, "%s :\t%#v\n", k, v)
	}
}
