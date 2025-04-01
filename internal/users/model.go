package users

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserModel struct {
	collection *mongo.Collection
}

func (userModel *UserModel) SetupModel(dbClient *mongo.Client) *UserModel {
	collection := dbClient.Database("fridge").Collection("users")
	userModel.collection = collection

	return userModel
}

func (user UserModel) CreateUser() {
	// user.collection.InsertOne(context.Background(), )
}
