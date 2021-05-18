package users

import (
	"context"
	"log"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ctx context.Context, input model.NewUser) (*model.User, error) {
	collection := database.GetCollection("users")

	// hash password before insert into database
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}
	res, err := collection.InsertOne(ctx, &UserType{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &model.User{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Name:     input.Name,
		Email:    input.Email,
		Password: "",
	}, nil
}

func FindById(ctx context.Context, id string) (*model.User, error) {
	collection := database.GetCollection("users")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	res := collection.FindOne(ctx, &UserType{
		ID: ObjectID,
	})
	user := UserType{}
	res.Decode(&user)
	return &model.User{
		ID:       user.ID.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: "",
	}, nil
}
