package repository

import (
	"context"
	"os"

	"github.com/danilomarques1/gos3example/api/database"
	"github.com/danilomarques1/gos3example/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(*model.User) error
	FindByEmail(string) (*model.User, error)
}

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() (UserRepository, error) {
	client, err := database.GetDbConnection()
	if err != nil {
		return nil, err
	}
	collection := client.Database(os.Getenv("DATABASE")).Collection("candidates")
	return &mongoUserRepository{collection: collection}, nil
}

func (m *mongoUserRepository) Save(user *model.User) error {
	if _, err := m.collection.InsertOne(context.Background(), user, options.InsertOne()); err != nil {
		return err
	}
	return nil
}

func (m *mongoUserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := m.collection.FindOne(context.Background(), bson.M{"email": email}, options.FindOne()).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
