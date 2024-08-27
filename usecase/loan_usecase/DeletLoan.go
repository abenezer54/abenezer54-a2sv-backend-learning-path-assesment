package loan_usecase

import (
	"context"
)

func (uc *loanUsecase) DeleteLoan(ctx context.Context, id string) error {
	return uc.loanRepo.DeleteLoan(ctx, id)
}
