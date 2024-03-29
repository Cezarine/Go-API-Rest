package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-type", "application/json")
		next.ServeHTTP(res, req)
	})
}
