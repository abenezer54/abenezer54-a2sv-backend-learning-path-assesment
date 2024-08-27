package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname          string             `json:"firstname" bson:"firstname"`
	Lastname           string             `json:"lastname" bson:"lastname"`
	Username           string             `json:"username" bson:"username"`
	Password           string             `json:"password" bson:"password"`
	Email              string             `json:"email" bson:"email"`
	Bio                string             `json:"bio" bson:"bio"`
	ProfilePicture     string             `json:"profile_picture" bson:"profile_picture"` // URL to the profile picture
	ContactInformation string             `json:"contact_information" bson:"contact_information"`
	IsAdmin            bool               `json:"isAdmin" bson:"isAdmin"`
	Active             bool               `json:"active" bson:"active"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
}

type UpdateRequest struct {
	Firstname          string `json:"firstname" bson:"firstname"`
	Lastname           string `json:"lastname" bson:"lastname"`
	Username           string `json:"username" bson:"username"`
	Bio                string `json:"bio" bson:"bio"`
	ProfilePicture     string `json:"profile_picture" bson:"profile_picture"`
	ContactInformation string `json:"contact_information" bson:"contact_information"`
	RefreshToken       string `json:"refresh_token" bson:"refresh_token"`
}

// type ForgotPasswordRequest struct {
// 	Email string `json:"email" binding:"required,email"`
// }

// type ResetPasswordRequest struct {
// 	ResetToken      string `json:"reset_token" binding:"required"`
// 	NewPassword     string `json:"new_password" binding:"required,min=8"`
// 	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
// }

type UserUsecase interface {

	// New methods for user usecase
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *UpdateRequest) error
	Register(ctx context.Context, req RegisterRequest, tokenSecret string, tokenExpiry int) (RegisterResponse, error)
	VerifyEmail(ctx context.Context, token string) error
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	// Logout(ctx context.Context, userID string) error
	RequestPasswordReset(ctx context.Context, email, frontendBaseURL string) error
	ResetPassword(ctx context.Context, req ResetPasswordRequest) error
	GetAllUsers(ctx context.Context) ([]User, error)
	DeleteUserByID(ctx context.Context, id string) error
	GetUserByID(ctx context.Context, id string) (*User, error)
	CreateUserFromToken(ctx context.Context, token string) (*User, error)
	// PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error
	// GetByEmail(ctx context.Context, email string) (User, error)
	// GetByUsername(ctx context.Context, username string) (User, error)
}

type UserRepository interface {
	// New methods for user repository

	CreateUser(ctx context.Context, user *User) error
	GetByUsernameOrEmail(ctx context.Context, identifier string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]User, error)
	// PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error
	// UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser *UpdateRequest) error
}
