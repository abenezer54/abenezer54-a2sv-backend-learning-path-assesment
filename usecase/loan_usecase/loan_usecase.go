package loan_usecase

import "loan-api/domain"

type loanUsecase struct {
	loanRepo domain.LoanRepository
}

func NewLoanUsecase(loanRepo domain.LoanRepository) *loanUsecase {
	return &loanUsecase{
		loanRepo: loanRepo,
	}
}
