package loan_repository

import (
	"context"
	"loan-api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *loanRepository) GetAllLoans(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	filter := bson.M{}
	if status != "" && status != "all" {
		filter["status"] = status
	}

	sortOrder := 1
	if order == "desc" {
		sortOrder = -1
	} else if order == "asc" {
		sortOrder = 1
	}

	findOptions := options.Find().SetSort(bson.D{{"status", sortOrder}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var loans []domain.Loan
	for cursor.Next(ctx) {
		var loan domain.Loan
		if err := cursor.Decode(&loan); err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}
