package loan_repository

import (
	"context"
	"loan-api/domain"
	"loan-api/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *loanRepository) GetLoanByID(ctx context.Context, id string) (*domain.LoanStatusResponse, error) {
	// Create a filter to match the loan by its ID
	filter := bson.M{"_id": id}

	// Define a variable to hold the result
	var loan domain.LoanStatusResponse

	// Find the loan in the collection
	err := r.collection.FindOne(ctx, filter).Decode(&loan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No matching loan found
			return nil, nil
		}
		// Return any other error
		return nil, err
	}

	// Return the found loan details
	return &loan, nil
}
