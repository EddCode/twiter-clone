package middlewares

import (
	"net/http"

	"github.com/EddCode/twitter-clone/internal/httpresponse"
	"github.com/EddCode/twitter-clone/utils"
)

func ValidJWT() middlewares {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			_, err := utils.ValidToken(r.Header.Get("Authorization"))
			if err != nil {
				httpresponse.Error("Unauthorized", err.Error()).Send(rw)
				return
			}

			next.ServeHTTP(rw, r)
		}
	}
}
