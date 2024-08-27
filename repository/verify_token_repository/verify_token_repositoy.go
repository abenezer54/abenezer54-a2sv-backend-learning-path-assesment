package verify_token_repository

import (
	"context"
	"errors"
	"loan-api/domain"
	"loan-api/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type verifyTokenRepository struct {
	collection mongo.Collection
}

func NewVerifyTokenRepository(collection mongo.Collection) domain.VerifyTokenRepository {
	return &verifyTokenRepository{
		collection: collection,
	}
}

func (r *verifyTokenRepository) ValidateVerifyToken(ctx context.Context, token string) (string, error) {
	var result struct {
		Email     string    `bson:"email"`
		ExpiresAt time.Time `bson:"expiresAt"`
	}

	filter := bson.M{"token": token}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("token not found")
		}
		return "", err
	}

	if time.Now().After(result.ExpiresAt) {
		return "", errors.New("token expired")
	}

	return result.Email, nil
}

func (r *verifyTokenRepository) InvalidateVerifyToken(ctx context.Context, token string) error {
	filter := bson.M{"token": token}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
