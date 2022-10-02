package data

import (
	"context"

	"yaosiqi525/cloud-platform/app/account/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) Save(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	err := r.data.DataBase.Save(g).Error
	return g, err
}

func (r *accountRepo) Update(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	return g, nil
}

func (r *accountRepo) FindByID(ctx context.Context, id int64) (info *biz.Account, err error) {
	err = r.data.DataBase.Where("id = ?", id).First(&info).Error
	return info, err
}

func (r *accountRepo) FindByAccount(ctx context.Context, account string) (info *biz.Account, err error) {
	err = r.data.DataBase.Where("account = ?", account).First(&info).Error
	return info, err
}

func (r *accountRepo) ListAll(context.Context) ([]*biz.Account, error) {
	return nil, nil
}
