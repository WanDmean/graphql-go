package users

import (
	"context"
	"fmt"

	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ctx context.Context, input Register) (*User, error) {
	/* hash password before insert into database */
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return &User{}, err
	}
	/* save input into database */
	res, err := database.Users.InsertOne(ctx, &User{
		Name:     input.Name,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Password: hashedPassword,
	})
	if err != nil {
		return &User{}, err
	}
	return &User{
		ID:     res.InsertedID.(primitive.ObjectID),
		Name:   input.Name,
		Email:  input.Email,
		Avatar: input.Avatar,
	}, nil
}

func FindById(ctx context.Context, id string) (*User, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &User{}, fmt.Errorf("invalid object id")
	}
	res := database.Users.FindOne(ctx, &User{
		ID: ObjectID,
	})
	user := User{}
	res.Decode(&user)
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func FindByEmail(ctx context.Context, email string) *User {
	res := database.Users.FindOne(ctx, &User{
		Email: email,
	})
	user := User{}
	res.Decode(&user)
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
