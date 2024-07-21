package middleware

import (
	"net/http"

	"github.com/raafly/ewallet/pkg/helper"
)

type AuthMiddleware struct {
	Handler http.Handler
}

type WebResponse struct {
	Code	int
	Status	string
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}
func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err != nil {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := WebResponse {
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		
		helper.WriteToResponBody(w, webResponse)
	} else {
		a.Handler.ServeHTTP(w, r)
	}
}