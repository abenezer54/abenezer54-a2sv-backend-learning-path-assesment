package verify_token_repository

import "loan-api/mongo"

type verifyTokenRepository struct {
	collection mongo.Collection
}

func NewVerifyTokenRepository(collection mongo.Collection) *verifyTokenRepository {
	return &verifyTokenRepository{
		collection: collection,
	}
}
