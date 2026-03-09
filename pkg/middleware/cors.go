package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return 
		}
		header := w.Header()
		header.Set("Acces-Control-Allow-Origin", origin)
		header.Set("Acces-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
        	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("ccess-Control-Max-Age", "86400")
		}
		next.ServeHTTP(w, r)
	})
}