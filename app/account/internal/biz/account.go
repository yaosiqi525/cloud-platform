package biz

import (
	"context"
	"crypto/md5"
	"fmt"

	v1 "yaosiqi525/cloud-platform/app/account/api/v1"
	"yaosiqi525/cloud-platform/pkg/random"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrAccountNotFound is account not found.
	ErrAccountNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "account not found")
	ErrAccountIsExist  = errors.BadRequest(v1.ErrorReason_USER_IS_EXIST.String(), "account is exist")
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
	FindByAccount(ctx context.Context, account string) (*Account, error)
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

	// 查找account是否存在
	_, err := uc.repo.FindByAccount(ctx, g.Account)
	if !(err != nil && IsNotFoundError(err)) {
		return nil, ErrAccountIsExist
	}

	// 密码简单加密
	g.Salt = random.GetRandomString()
	g.Pwd = g.Salt + g.Pwd
	g.Pwd = fmt.Sprintf("%x", md5.Sum([]byte(g.Pwd)))

	// 保存
	return uc.repo.Save(ctx, g)
}
