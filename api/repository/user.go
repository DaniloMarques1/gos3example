package repository

import (
	"context"
	"errors"
	"os"

	"github.com/danilomarques1/gos3example/api/database"
	"github.com/danilomarques1/gos3example/api/model"
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
	return nil, errors.New("To be implemented")
}
