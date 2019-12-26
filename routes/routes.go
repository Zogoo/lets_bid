package routes

import (
	v1 "lets_bid/api/v1"
	"lets_bid/service"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	r.Use(TokenValidator)

	r.HandleFunc("/cas/login", v1.Auth).Methods("POST")
	r.HandleFunc("/cas/register", v1.Create).Methods("POST")
	r.HandleFunc("/sso/my_page", v1.Index).Methods("GET")

	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Add("Vary", "Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, X-Requested-With")
		next.ServeHTTP(w, r)
	})
}

// TokenValidator validate header
func TokenValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		authorizationHeader := r.Header.Get("Authorization")
		jwtTokenString := strings.Split(authorizationHeader, " ")[1]

		if jwtTokenString != "" && jwtTokenString != "null" {
			_, err = service.ValidateToken(jwtTokenString)
		}
		if err != nil {
			http.Error(w, "401 Unauthorized request", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}

		defer func() {
			if r := recover(); r != nil {
				http.Error(w, "504 Internal service error", http.StatusInternalServerError)
			}
		}()
	})
}
