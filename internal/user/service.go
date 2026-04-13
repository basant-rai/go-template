package user

import (
	"context"

	"github.com/go-template/config"
)

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, rawRefreshToken string) (*RefreshTokenResponse, error)
	Logout(ctx context.Context, rawRefreshToken string) error
}

type service struct {
	repo   Repository
	config *config.Config
}

func NewService(repo Repository, cfg *config.Config) Service {
	return &service{repo: repo, config: cfg}
}
