package loan_usecase

import (
	"context"
	"errors"
	"loan-api/domain"
)

func (uc *loanUsecase) UpdateLoanStatus(ctx context.Context, id string, status string) (*domain.Loan, error) {
	if status != "approved" && status != "rejected" {
		return nil, errors.New("invalid status")
	}

	return uc.loanRepo.UpdateLoanStatus(ctx, id, status)
}
