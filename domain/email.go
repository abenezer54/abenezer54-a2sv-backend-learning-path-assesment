package domain

import "context"

type EmailService interface {
	SendPasswordResetEmail(ctx context.Context, email, resetLink string) error
	SendVerificationEmail(ctx context.Context, email, resetLink string) error
}
