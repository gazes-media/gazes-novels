package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"gorm.io/gorm"
)

const AuthUrl string = "https://api.gazes.fr/auth"

func AuthGetUserMe(token string) (*UserMiddleware, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", AuthUrl+"/@me", nil)
	if err != nil {
		return nil, err
	}

	// send the token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	var user UserMiddleware
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AuthPostLogin(email, password string) (*string, error) {
	bodyReader := BuildIoReaderFromMap(map[string]interface{}{
		"email":    email,
		"password": password,
	})
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", AuthUrl+"/login", bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	var token string
	if res.Header.Get("Set-Cookie") == "" {
		return nil, err
	} else {
		cookie := cookieHeader(res.Header.Get("Set-Cookie"))
		token = cookie[0].Value
	}

	return &token, nil
}

func AuthPostRegister(username, email, password string) (*string, error) {
	bodyReader := BuildIoReaderFromMap(map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
	})
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", AuthUrl+"/register", bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	var token string

	if res.Header.Get("Set-Cookie") == "" {
		return nil, err
	} else {
		cookie := cookieHeader(res.Header.Get("Set-Cookie"))
		token = cookie[0].Value
	}

	return &token, nil
}

func BuildIoReaderFromMap(data map[string]interface{}) io.Reader {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return bytes.NewReader(jsonData)
}

func cookieHeader(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}
	return req.Cookies()
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthdate time.Time `json:"birthdate"`
}

type UserMiddleware struct {
	gorm.Model
	Username string `json:"username"`            // Username is the user's username.
	Email    string `gorm:"unique" json:"email"` // Email is the user's email address.
	Password string `json:"password"`            // Password is the user's password.
}
