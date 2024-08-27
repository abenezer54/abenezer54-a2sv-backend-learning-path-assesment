package loan_usecase

import (
	"context"
	"loan-api/domain"
)

func (uc *loanUsecase) GetLoanStatus(ctx context.Context, id string) (*domain.LoanStatusResponse, error) {
	// Fetch loan status from the repository
	loanStatus, err := uc.loanRepo.GetLoanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return loanStatus, nil
}
