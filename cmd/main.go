package main

import (
	"github.com/nazobenkyou/gorm-manabu/config"
	"github.com/nazobenkyou/gorm-manabu/input/http_router"
	"net/http"
)

func main() {
	config.LoadConfig()

	panic(http.ListenAndServe(config.AppConfig.SrvPort, http_router.NewRouter()))
}
