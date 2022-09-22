package biz

import (
	"context"

	v1 "yaosiqi525/cloud-platform/app/account/api/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrAccountNotFound is account not found.
	ErrAccountNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "account not found")
)

// Account is a Account model.
type Account struct {
	Account string
	Name    string
	Pwd     string
	Salt    string
	Status  int32
	OrmModel
}

// AccountRepo is a Greater repo.
type AccountRepo interface {
	Save(context.Context, *Account) (*Account, error)
	Update(context.Context, *Account) (*Account, error)
	FindByID(context.Context, int64) (*Account, error)
	ListByHello(context.Context, string) ([]*Account, error)
	ListAll(context.Context) ([]*Account, error)
}

// AccountUsecase is a Account usecase.
type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

// NewAccountUsecase new a Account usecase.
func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAccount creates a Account, and returns the new Account.
func (uc *AccountUsecase) CreateAccount(ctx context.Context, g *Account) (*Account, error) {
	uc.log.WithContext(ctx).Infof("CreateAccount: %v", g.Account)
	return uc.repo.Save(ctx, g)
}
