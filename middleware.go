package main

import (
	"log"
	"net/http"
)

func adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

type middleware func(http.HandlerFunc) http.HandlerFunc

func adapt(f http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func logRequest() middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
			next.ServeHTTP(w, r)
		}
	}
}

func secureHeaders() middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "deny")
			next.ServeHTTP(w, r)
		}
	}
}
