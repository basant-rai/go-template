package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	// ExistsByEmail(ctx context.Context, email string) (bool, error)
	// Create(ctx context.Context, u *User) error
	// GetByEmail(ctx context.Context, email string) (*User, error)
	// GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	// UpdateStripeAccount(ctx context.Context, id uuid.UUID, stripeAccountID string) error
	// UpdateKYB(ctx context.Context, id uuid.UUID, result KYBResult, req KYBRequest) error
	// UpdateBankAccount(ctx context.Context, id uuid.UUID, routing, account, last4 string) error
	// GetStripeAccountID(ctx context.Context, id uuid.UUID) (string, error)
	// ResetMPINAttempts(ctx context.Context, email string) error
	// IncrementMPINAttempts(ctx context.Context, email string) error
	// LockMPIN(ctx context.Context, email string, until time.Time) error
	// // Refresh tokens
	// SaveRefreshToken(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error
	// GetRefreshToken(ctx context.Context, tokenHash string) (userID uuid.UUID, revoked bool, err error)
	// RevokeRefreshToken(ctx context.Context, tokenHash string) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}
