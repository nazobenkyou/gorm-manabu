package main

import (
	"github.com/nazobenkyou/gorm-manabu/config"
	"github.com/nazobenkyou/gorm-manabu/input/http_router"
	"github.com/nazobenkyou/gorm-manabu/repository"
	"net/http"
)

func main() {
	config.LoadConfig()

	conn, err := repository.Connect()
	if err != nil {
		panic(err)
	}

	defer conn.Conn.Close()

	conn.Conn.AutoMigrate(&repository.Task{})

	panic(http.ListenAndServe(config.AppConfig.SrvPort, http_router.NewRouter(conn)))
}
