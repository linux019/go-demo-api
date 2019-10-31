package controllers

import (
	"api-demo/api/util"
	"api-demo/apilogger"
	"api-demo/db"
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := decodeJSON(&user, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apilogger.Logger.Error("CreateAccount", err)
	} else {
		if user.Password != "" && strings.TrimSpace(user.Email) != "" {
			hashedPwd, _ := util.GenerateFromPassword(user.Password)
			result := db.CreateUser(user.Email, hashedPwd)

			if result {
				apilogger.Logger.Info("Created user ", user.Email)
				w.WriteHeader(http.StatusOK)
			} else {
				apilogger.Logger.Info("Skip creation duplicate of user ", user.Email)
				w.WriteHeader(http.StatusCreated)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func AuthenticateAccount(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := decodeJSON(&user, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apilogger.Logger.Error("AuthenticateAccount", err)
	} else {
		if user.Password != "" && strings.TrimSpace(user.Email) != "" {

			if result, ok := db.GetUser(user.Email); ok {
				if util.ComparePasswordAndHash(user.Password, result.Passwd) {
					apilogger.Logger.Info("AUTH ", user.Email, " valid password")
					w.WriteHeader(http.StatusOK)
				} else {
					apilogger.Logger.Info("AUTH ", user.Email, " invalid password")
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				apilogger.Logger.Info("AUTH ", user.Email, "not found")
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func decodeJSON(v interface{}, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(v)
	return err
}

func Demo(w http.ResponseWriter, r *http.Request) {

}
