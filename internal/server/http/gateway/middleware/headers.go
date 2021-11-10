package middleware

import "net/http"

func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Service-Name", "Gateway-Service")
		next.ServeHTTP(w, r)
	})
}
