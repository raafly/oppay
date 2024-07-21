package middleware

import (
	"net/http"

)

type LogMiddleware struct {
	Handler http.Handler
}

func (l *LogMiddleware) ServerHttp(w http.ResponseWriter, r *http.Request) {
	l.Handler.ServeHTTP(w,r)
}