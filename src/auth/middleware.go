package auth

import (
	"context"
	"net/http"

	"github.com/WanDmean/graphql-go/src/pkg/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// validate jwt token
			tokenStr := header
			userId, err := users.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			// create user and check if user exists in db
			user, err := users.FindById(r.Context(), userId)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.UserType {
	raw, _ := ctx.Value(userCtxKey).(*users.UserType)
	return raw
}
