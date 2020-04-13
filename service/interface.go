package service

import (
	"github.com/nazobenkyou/gorm-manabu/model"
	"github.com/nazobenkyou/gorm-manabu/repository"
)

type (
	TaskService struct {
		Repository repository.Repository
	}

	TaskServicer interface {
		GetAll() ([]model.Task, error)
	}
)
