package config

type (
	App struct {
		LogLevel string `default:"INFO"`
		SrvPort  string `default:":8080"`
	}
)

var (
	AppConfig App
)
