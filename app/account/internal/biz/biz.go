package biz

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAccountUsecase)

type OrmModel struct {
	ID uint64 `gorm:"primary_key;autoIncrement;"`
	OrmSimpleModel
}

type OrmSimpleModel struct {
	CreatedAt int64                 `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt int64                 `gorm:"column:update_time;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;comment:删除时间" json:"-"`
}

type OrmUUIDModel struct {
	ID string `gorm:"primary_key"`
	OrmSimpleModel
}

func IsNotFoundError(err error) bool {
	return gorm.ErrRecordNotFound.Error() == err.Error()
}
