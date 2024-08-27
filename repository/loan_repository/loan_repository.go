package loan_repository

import (
	"loan-api/mongo"
)

type loanRepository struct {
	collection mongo.Collection
}

func NewLoanRepository(collection mongo.Collection) *loanRepository {
	return &loanRepository{
		collection: collection,
	}
}
