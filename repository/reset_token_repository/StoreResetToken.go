package reset_token_repository

import (
	"loan-api/domain"
	"context"
)

func (r *resetTokenRepository) StoreResetToken(ctx context.Context, token domain.PasswordResetToken) error {
	_, err := r.collection.InsertOne(ctx, token)
	return err
}
