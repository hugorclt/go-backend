package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//do some validation with db

			ctx := context.WithValue(r.Context(), ContextKey("userId"), userId)
			next.ServeHttp(w, r.WithContext(ctx))
		})
	}
}

func getUserIdFromContext(r *http.Request) interface{} {
	return r.Context().Value(ContextKey("userId"))
}
