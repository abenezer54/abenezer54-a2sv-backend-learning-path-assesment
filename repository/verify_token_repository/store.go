package verify_token_repository

import (
	"context"
)

func (r *verifyTokenRepository) StoreVerifyToken(ctx context.Context, token string) error {
	_, err := r.collection.InsertOne(ctx, token)
	return err

}
