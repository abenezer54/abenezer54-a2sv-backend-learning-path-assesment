package loan_usecase

import (
	"context"
	"loan-api/domain"
)

func (uc *loanUsecase) ApplyForLoan(ctx context.Context, req *domain.LoanApplicationRequest) (*domain.LoanApplicationResponse, error) {
	loan := &domain.Loan{
		UserID:  req.UserID,
		Amount:  req.Amount,
		Term:    req.Term,
		Purpose: req.Purpose,
		Status:  "submitted",
	}

	// Store the loan application in the database
	storedLoan, err := uc.loanRepo.StoreLoan(ctx, loan)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := &domain.LoanApplicationResponse{
		ApplicationID: storedLoan.ApplicationID.Hex(), // Convert ObjectID to hex string
		Status:        storedLoan.Status,
		Message:       "Your loan application has been submitted successfully.",
	}

	return response, nil
}
