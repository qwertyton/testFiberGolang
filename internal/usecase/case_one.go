package usecase

import (
	"goTestFiber/model"
	"goTestFiber/repository"
	"gorm.io/gorm"
)

func SumFromDB(db *gorm.DB) model.Respond {
	repo := repository.InitTestRepo(db)
	return repo.Sum()
}
