package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ApplicationID primitive.ObjectID `bson:"_id,omitempty" json:"application_id"`
	UserID        string             `bson:"user_id" json:"user_id"`
	Amount        float64            `bson:"amount" json:"amount"`
	Term          int                `bson:"term" json:"term"`
	Purpose       string             `bson:"purpose" json:"purpose"`
	Status        string             `bson:"status" json:"status"`
}

type LoanApplicationRequest struct {
	UserID  string  `json:"user_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
	Term    int     `json:"term" binding:"required"` // in months
	Purpose string  `json:"purpose" binding:"required"`
}

type LoanApplicationResponse struct {
	ApplicationID string `json:"application_id"`
	Status        string `json:"status"`
	Message       string `json:"message,omitempty"`
}

type LoanUsecase interface {
	ApplyForLoan(ctx context.Context, req *LoanApplicationRequest) (*LoanApplicationResponse, error)
	GetLoanStatus(ctx context.Context, id string) (*LoanStatusResponse, error)
	GetAllLoans(ctx context.Context, status string, order string) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, id string, status string) (*Loan, error)
	DeleteLoan(ctx context.Context, id string) error
}
type LoanRepository interface {
	StoreLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, id string) (*LoanStatusResponse, error)
	GetAllLoans(ctx context.Context, status string, order string) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, id string, status string) (*Loan, error)
	DeleteLoan(ctx context.Context, id string) error
}

type LoanStatusResponse struct {
	ApplicationID string  `json:"application_id"`
	Status        string  `json:"status"`
	Amount        float64 `json:"amount"`
	Term          int     `json:"term"` // in months
	Purpose       string  `json:"purpose"`
	UserID        string  `json:"user_id"`
}

type UpdateLoanStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
