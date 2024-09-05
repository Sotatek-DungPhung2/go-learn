package main

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (app *application) authHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		app.writeJSON(w, http.StatusBadRequest, "Username or password can not be empty", nil)
		return
	}

	httpClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3003/data?key=dung", nil)

	if err != nil {
		app.writeJSON(w, http.StatusBadRequest, "dcmm sai r", nil)
		return
	}

	res, err := httpClient.Do(req)

	if err != nil {
		app.writeJSON(w, http.StatusBadRequest, "dcmm sai r 1", nil)
		return
	}

	var users []User

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&users)

	if err != nil {
		app.writeJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if len(users) > 0 {
		if username == users[0].Username && password == users[0].Password {
			app.writeJSON(w, http.StatusOK, "ok", nil)
		} else {
			app.writeJSON(w, http.StatusUnauthorized, "dm", nil)
		}
	}
}
