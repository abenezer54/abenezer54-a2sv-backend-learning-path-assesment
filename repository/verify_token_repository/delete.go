package verify_token_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *verifyTokenRepository) DeleteVerifyToken(ctx context.Context, userID string) error {
	filter := bson.M{"user_id": userID}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
