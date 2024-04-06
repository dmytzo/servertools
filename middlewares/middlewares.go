package middlewares

import (
	"log"
	"net/http"
)

func RequestLogMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[request] %s %s", r.Method, r.RequestURI)

		handlerFunc(w, r)
	}
}
