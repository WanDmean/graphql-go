package users

import (
	"context"
	"fmt"

	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ctx context.Context, input Register) (*UserType, error) {
	/* hash password before insert into database */
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return &UserType{}, err
	}
	/* save input into database */
	res, err := database.Users.InsertOne(ctx, &UserType{
		Name:     input.Name,
		Email:    input.Email,
		Avatar:   input.Avatar,
		Password: hashedPassword,
	})
	if err != nil {
		return &UserType{}, err
	}
	return &UserType{
		ID:     res.InsertedID.(primitive.ObjectID),
		Name:   input.Name,
		Email:  input.Email,
		Avatar: input.Avatar,
	}, nil
}

func FindById(ctx context.Context, id string) (*UserType, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &UserType{}, fmt.Errorf("invalid object id")
	}
	res := database.Users.FindOne(ctx, &UserType{
		ID: ObjectID,
	})
	user := UserType{}
	res.Decode(&user)
	return &UserType{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func FindByEmail(ctx context.Context, email string) *UserType {
	res := database.Users.FindOne(ctx, &UserType{
		Email: email,
	})
	user := UserType{}
	res.Decode(&user)
	return &UserType{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
