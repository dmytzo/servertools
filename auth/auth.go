package auth

import (
	"context"
	"net/http"
	"time"
)

const (
	userCtxKey      = "user_id"
	sessionTokenKey = "session_token"
	sessionTokenTTL = time.Hour * 24 * 365
)

type User struct{ ID, Username string }

func UserFromRequest(r *http.Request) (User, bool) {
	u := r.Context().Value(userCtxKey)
	if u == nil {
		return User{}, false
	}

	user, ok := u.(User)
	if !ok {
		return User{}, false
	}

	return user, true
}

func UserFromRequestOrEmpty(r *http.Request) User {
	user, _ := UserFromRequest(r)
	return user
}

func RequestWithUser(r *http.Request, user User) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userCtxKey, user))
}

func IsAuthorized(r *http.Request) bool {
	_, exists := UserFromRequest(r)
	return exists
}

func SetSessionTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionTokenKey,
		Expires: time.Now().Add(sessionTokenTTL),
		Value:   token,
	})
}

func SessionTokenCookieFromRequest(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(sessionTokenKey)
}

func DeleteSessionTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  sessionTokenKey,
		Value: "",
	})
}
