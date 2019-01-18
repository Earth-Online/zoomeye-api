package zoomeye // github.com/Earth-Online/zoomeye-api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//  Zoomeye login get token
func UserLogin(username string, password string) (Token, error) {
	var login = Login{
		Username: username,
		Password: password,
	}
	var jsonStr, _ = json.Marshal(login)
	resp, err := http.Post(LoginApi, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return Token{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return Token{}, errors.New(string(body))
	}
	token := Token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return Token{}, errors.New(string(body))
	}
	return token, nil
}

// Get new user
func NewUser(username string, password string) (User, error) {
	token, err := UserLogin(username, password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Login{
			username,
			password,
		},
		token,
	}
	return user, nil
}
