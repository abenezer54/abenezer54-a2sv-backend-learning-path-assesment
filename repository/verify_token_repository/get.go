package verify_token_repository

import (
	"context"
	"loan-api/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *verifyTokenRepository) GetVerifyToken(ctx context.Context, userID string) (string, error) {
	filter := bson.M{"user_id": userID}
	var result struct {
		VerifyToken string `bson:"verify_token"`
	}

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil // No token found for this user
		}
		return "", err
	}

	return result.VerifyToken, nil
}
