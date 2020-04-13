package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nazobenkyou/gorm-manabu/config"
)

type (
	MariaDBRepository struct {
		Conn *gorm.DB
	}
)

func Connect() (*MariaDBRepository, error) {
	if conn, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=%s", config.Database.User, config.Database.Password, config.Database.DatabaseName, config.Database.Charset)); err != nil {
		return nil, err
	} else {
		return &MariaDBRepository{Conn: conn}, nil
	}
}

func (db *MariaDBRepository) GetAll() []Task {
	var tasks []Task

	db.Conn.Find(&tasks)

	return tasks
}
