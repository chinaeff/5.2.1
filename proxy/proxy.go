package proxy

import "net/http"

func ProxyHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthorized(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isAuthorized(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	return token == "some token"
}
