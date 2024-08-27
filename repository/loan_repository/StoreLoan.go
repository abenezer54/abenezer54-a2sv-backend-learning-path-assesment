package loan_repository

import (
	"context"
	"loan-api/domain"
)

func (r *loanRepository) StoreLoan(ctx context.Context, loan *domain.Loan) (*domain.Loan, error) {
	// Insert the loan into the collection
	_, err := r.collection.InsertOne(ctx, loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}
