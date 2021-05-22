package auth

import (
	"encoding/json"
	"net/http"

	"github.com/WanDmean/graphql-go/src/pkg/users"
	"github.com/WanDmean/graphql-go/src/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	/* decode request body */
	var register users.Register
	json.NewDecoder(r.Body).Decode(&register)
	/* check if email is exist */
	user := users.FindByEmail(r.Context(), register.Email)
	if user.Email != "" {
		/* throw bad request error */
		util.ResponseError(w, "email already exist", 400)
		return
	}
	/* save input into database */
	user, err := users.Save(r.Context(), register)
	if err != nil {
		/* throw internal server error */
		util.ResponseError(w, err.Error(), 500)
		return
	}
	/* generate Token */
	token, _ := GenerateToken(user.ID.Hex())
	/* response json */
	util.ResponseJson(w, token, 201)
}

func Login(w http.ResponseWriter, r *http.Request) {
	/* decode request body */
	var login users.Login
	json.NewDecoder(r.Body).Decode(&login)
	/* check if email is exist */
	user := users.FindByEmail(r.Context(), login.Email)
	if user.Email == "" {
		/* throw bad request error */
		util.ResponseError(w, "user does not exist", 400)
		return
	}
	if users.CheckPasswordHash(login.Password, user.Password) {
		/* generate Token */
		token, _ := GenerateToken(user.ID.Hex())
		/* response json */
		util.ResponseJson(w, token, 201)
		return
	}
	/* throw unauthorized error */
	util.ResponseError(w, "invalid password", 401)
}
