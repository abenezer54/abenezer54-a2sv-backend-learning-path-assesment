package domain

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"` // Ensure this is hashed and securely handled
	jwt.StandardClaims
}

type VerificationToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Token     string             `bson:"token"`
	ExpiresAt time.Time          `bson:"expires_at"`
	CreatedAt time.Time          `bson:"created_at"`
}

type RefreshTokenRepository interface {
	StoreRefreshToken(ctx context.Context, userID string, tokenString string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, userID string) (string, error)
	DeleteRefreshToken(ctx context.Context, userID string) error
}

type PasswordResetToken struct {
	Token  string    `json:"token"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
	Used   bool      `json:"used"`
}

type ResetTokenRepository interface {
	StoreResetToken(ctx context.Context, token PasswordResetToken) error
	ValidateResetToken(ctx context.Context, token string) (string, error)
	InvalidateResetToken(ctx context.Context, token string) error
}
type VerifyTokenRepository interface {
	StoreVerifyToken(ctx context.Context, token string, email string) error
	ValidateVerifyToken(ctx context.Context, token string) (string, error)
	InvalidateVerifyToken(ctx context.Context, token string) error
}
