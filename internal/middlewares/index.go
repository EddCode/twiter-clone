package middlewares

import "net/http"

type middlewares func(http.HandlerFunc) http.HandlerFunc

func AttachMiddelwares(handler http.HandlerFunc, middlewares ...middlewares) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
