package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/ewallet/internal/core"
	"github.com/raafly/ewallet/pkg/helper"
	"github.com/raafly/ewallet/pkg/middleware"
)

func NewRoute(user core.UserHandler) *httprouter.Router {
	route := httprouter.New()

	route.POST("/api/signup", user.SignUp)
	route.POST("/api/singin", user.SignIn)
	route.GET("/api/users/:id", user.FindById)
	// route.GET("/api/protected/users/", protected(user.SignUp))
	route.PanicHandler = core.ErrorHandler
	return route
}

func protected(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if _, err := r.Cookie("auth"); err != nil {
			// error
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			webResponse := middleware.WebResponse {
				Code: http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponBody(w, webResponse)
		}
	}
}