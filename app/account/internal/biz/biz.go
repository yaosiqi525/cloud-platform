package biz

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAccountUsecase)

type OrmModel struct {
	ID uint64 `gorm:"primary_key;autoIncrement;"`
	OrmSimpleModel
}

type OrmSimpleModel struct {
	CreatedAt int64 `gorm:"column:create_time;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:update_time;autoUpdateTime"`
}

type OrmUUIDModel struct {
	ID        string         `gorm:"primary_key"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:delete_time"`

	OrmSimpleModel
}
