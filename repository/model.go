package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nazobenkyou/gorm-manabu/model"
)

type (
	Task struct {
		gorm.Model

		model.Task
	}
)
