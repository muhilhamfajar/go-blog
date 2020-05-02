package middlewares

import "net/http"

func SetMiddlewareJson(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next(writer, r)
	}
}