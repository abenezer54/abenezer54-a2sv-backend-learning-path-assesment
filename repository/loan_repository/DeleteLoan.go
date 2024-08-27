package loan_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *loanRepository) DeleteLoan(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid loan ID")
	}

	filter := bson.M{"_id": objID}
	deletedCount, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if deletedCount == 0 {
		return errors.New("loan application not found")
	}

	return nil
}
