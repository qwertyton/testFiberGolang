package repository

import (
	"goTestFiber/model"
	"gorm.io/gorm"
)

type TestRepository interface {
	Sum() model.Respond
}

type testRepo struct {
	db *gorm.DB
}

func InitTestRepo(db *gorm.DB) TestRepository {
	return &testRepo{db: db}
}
