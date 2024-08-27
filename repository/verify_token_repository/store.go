package verify_token_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *verifyTokenRepository) StoreVerifyToken(ctx context.Context, token string, email string) error {
	_, err := r.collection.InsertOne(ctx, bson.M{
		"token": token,
		"email": email,
	})
	if err != nil {
		return err
	}
	return nil
}
