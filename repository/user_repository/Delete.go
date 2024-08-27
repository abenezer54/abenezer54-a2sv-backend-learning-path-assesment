package user_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
