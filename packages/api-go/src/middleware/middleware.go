package middlewares

import (
	"log"
	"net/http"

	"github.com/otaviopontes/api-go/src/authentication"
	"github.com/otaviopontes/api-go/src/responses"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
