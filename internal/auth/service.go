package auth

import (
	"context"

	"github.com/go-template/config"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	LoginMPIN(ctx context.Context, req LoginMPINRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, rawRefreshToken string) (*RefreshTokenResponse, error)
	Logout(ctx context.Context, rawRefreshToken string) error
	SubmitKYB(ctx context.Context, userID string, req KYBRequest) (*KYBResponse, error)
	AddBankAccount(ctx context.Context, userID string, req BankAccountRequest) error
	SearchUser(ctx context.Context, email string) (*UserSearchResult, error)
	GetMe(ctx context.Context, userID string) (*User, error)
	GetAccountStatus(ctx context.Context, userID string) (map[string]interface{}, error)
}

type service struct {
	repo   Repository
	stripe *stripeClient.Client
	config *config.Config
}

func NewService(repo Repository, stripe *stripeClient.Client, cfg *config.Config) Service {
	return &service{repo: repo, stripe: stripe, config: cfg}
}
