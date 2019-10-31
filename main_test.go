package main

import (
	"api-demo/api/controllers"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

const serverAddr = "http://127.0.0.1"

var httpClient http.Client = http.Client{
	Timeout: time.Second * 5,
}

func serverRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	if req, err := http.NewRequest(method, serverAddr+url, body); err == nil {
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Content-Type", "application/json")
		return req
	} else {
		t.Fail()
	}

	return nil
}

func TestServer(t *testing.T) {
	t.Run("Connection test", func(t *testing.T) {
		req := serverRequest(t, http.MethodHead, "/", nil)
		if r, err := httpClient.Do(req); err == nil {
			defer r.Body.Close()
			if r.StatusCode != http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v want %v",
					r.StatusCode, http.StatusOK)
			}
		} else {
			t.Error(err)
		}
	})

	b := make([]byte, 8, 8)
	var err error

	if _, err = rand.Read(b); err != nil {
		t.Fail()
	}

	email := "test_" + hex.EncodeToString(b)
	passwd := hex.EncodeToString(b)
	jsonStr, _ := json.Marshal(controllers.User{Email: email, Password: passwd})

	t.Run("Create user test", func(t *testing.T) {
		req := serverRequest(t, http.MethodPost, "/create_account", bytes.NewBuffer(jsonStr))

		if r, err := httpClient.Do(req); err == nil {
			defer r.Body.Close()
			if r.StatusCode != http.StatusOK && r.StatusCode != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v want %v",
					r.StatusCode, http.StatusOK)
			}
		} else {
			t.Error(err)
		}
	})

	t.Run("Authorization test", func(t *testing.T) {
		req := serverRequest(t, http.MethodPost, "/authenticate", bytes.NewBuffer(jsonStr))

		if r, err := httpClient.Do(req); err == nil {
			defer r.Body.Close()
			if r.StatusCode != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					r.StatusCode, http.StatusOK)
			}
		} else {
			t.Error(err)
		}
	})

	if _, err = rand.Read(b); err != nil {
		t.Fail()
	}

	email = "test_" + hex.EncodeToString(b)
	passwd = hex.EncodeToString(b)
	jsonStr, _ = json.Marshal(controllers.User{Email: email, Password: passwd})

	t.Run("Authorization test(incorrect user)", func(t *testing.T) {
		req := serverRequest(t, http.MethodPost, "/authenticate", bytes.NewBuffer(jsonStr))

		if r, err := httpClient.Do(req); err == nil {
			defer r.Body.Close()
			if r.StatusCode != http.StatusUnauthorized {
				t.Errorf("handler returned wrong status code: got %v want %v",
					r.StatusCode, http.StatusOK)
			}
		} else {
			t.Error(err)
		}
	})

}
