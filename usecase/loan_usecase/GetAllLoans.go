package loan_usecase

import (
	"context"
	"loan-api/domain"
)

func (uc *loanUsecase) GetAllLoans(ctx context.Context, status string, order string) ([]domain.Loan, error) {
	return uc.loanRepo.GetAllLoans(ctx, status, order)
}
