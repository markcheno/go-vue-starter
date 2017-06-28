package auth

import (
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/markcheno/go-vue-starter/models"
)

// signingKey set up a global string for our secret
var signingKey = []byte("knrjkevdckjh")

// JwtMiddleware handler for jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	},
	UserProperty:  "user",
	Debug:         false,
	SigningMethod: jwt.SigningMethodHS256,
})

// GetToken create a jwt token with user claims
func GetToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}

// GetJSONToken create a JSON token string
func GetJSONToken(user *models.User) string {
	token := GetToken(user)
	jsontoken := "{\"id_token\": \"" + token + "\"}"
	return jsontoken
}

// GetUserClaimsFromContext return "user" claims as a map from request
func GetUserClaimsFromContext(req *http.Request) map[string]interface{} {
	//claims := context.Get(req, "user").(*jwt.Token).Claims.(jwt.MapClaims)
	claims := req.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return claims
}
