package api

import (
	"encoding/json"
	"net/http"

	"github.com/markcheno/go-vue-starter/auth"
	"github.com/markcheno/go-vue-starter/models"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserSignup -
func (api *API) UserSignup(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(jsondata.Username) {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(jsondata.Username, jsondata.Password)

	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

// UserLogin -
func (api *API) UserLogin(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	user := api.users.FindUser(jsondata.Username)
	if user.Username == "" {
		http.Error(w, "username not found", http.StatusBadRequest)
		return
	}

	if !api.users.CheckPassword(user.Password, jsondata.Password) {
		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))

}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
