package loan_repository

import (
	"context"
	"errors"
	"loan-api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *loanRepository) UpdateLoanStatus(ctx context.Context, id string, status string) (*domain.Loan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid loan ID")
	}

	// Filter to find the loan by its ID
	filter := bson.M{"_id": objID}

	// Update the loan status
	update := bson.M{"$set": bson.M{"status": status}}
	updateResult, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Check if the loan was found and updated
	if updateResult.MatchedCount == 0 {
		return nil, errors.New("loan application not found")
	}

	// Retrieve the updated loan document
	var updatedLoan domain.Loan
	err = r.collection.FindOne(ctx, filter).Decode(&updatedLoan)
	if err != nil {
		return nil, err
	}

	return &updatedLoan, nil
}
