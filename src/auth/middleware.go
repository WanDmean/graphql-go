package auth

import (
	"context"
	"html"
	"net/http"
	"strings"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/pkg/users"
	"github.com/WanDmean/graphql-go/src/util"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow unauthenticated path in
			if UnauthenticatedPath(r.Method + html.EscapeString(r.URL.Path)) {
				next.ServeHTTP(w, r)
				return
			}
			// validate jwt token
			tokenStr := r.Header.Get("Authorization")
			userId, err := ParseToken(tokenStr)
			if err != nil {
				util.ResponseError(w, "invalid token", 401)
				return
			}
			// create user and check if user exists in db
			user, err := users.FindById(r.Context(), userId)
			if err != nil {
				util.ResponseError(w, "not found user", 404)
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
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

// Check unauthenticated path
func UnauthenticatedPath(inputPath string) bool {
	unauthPath := []string{
		"post/api/register",
		"post/api/login",
	}
	for _, path := range unauthPath {
		if path == strings.ToLower(inputPath) {
			return true
		}
	}
	return false
}
