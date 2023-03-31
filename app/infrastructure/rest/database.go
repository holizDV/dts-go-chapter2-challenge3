package rest

import (
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/entity"

	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/config"
	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitDatabase(dbConfig postgres.Config) *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}
	dsn, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	helper.PanicIfError(err)

	if config.IsMigrate && dbInstance == nil {
		err := dsn.AutoMigrate(&entity.Book{})
		helper.PanicIfError(err)

	}
	dbInstance = dsn
	return dbInstance
}
