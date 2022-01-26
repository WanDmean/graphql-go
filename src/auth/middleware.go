package auth

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/pkg/users"
	"github.com/WanDmean/graphql-go/src/util"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
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

		ctx := r.Context()
		// create user and check if user exists in db
		user, err := users.FindById(ctx, userId)
		if err != nil {
			util.ResponseError(w, "not found user", 404)
			return
		}
		// put it in context
		ctx = context.WithValue(ctx, UserCtxKey, &model.User{
			ID:     user.ID.Hex(),
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		})
		// and call the next with our new context
		fmt.Println(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	// NOTE if not found user
	// may be b'coz when set user in context
	// user type is not [*model.User]
	raw, _ := ctx.Value(UserCtxKey).(*model.User)
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
