package user_repository

import (
	"loan-api/mongo"
)

type userRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) *userRepository {
	return &userRepository{
		collection: collection,
	}
}
