package auth

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"video-processing/model"
	"video-processing/services"
)

func AuthUser(w http.ResponseWriter, r *http.Request) (model.User, error) {
	basicAuthPrefix := "Basic "
	authStr := r.Header.Get("Authorization")
	var user model.User
	if strings.HasPrefix(authStr, basicAuthPrefix) {
		payload, err := base64.StdEncoding.DecodeString(
			authStr[len(basicAuthPrefix):],
		)
		if err == nil {
			pair := bytes.SplitN(payload, []byte(":"), 2)
			if len(pair) == 2 {
				account := string(pair[0])
				password := string(pair[1])
				us := services.GetUserService()
				user, err = us.GetUserByEmail(account)
				if err != nil {
					return user, err
				}
				if user.Password == password {
					return user, nil
				} else {
					return user, errors.New("User password wrong")
				}
			}
		}
	}

	return user, errors.New("User basic auth failed")
}
