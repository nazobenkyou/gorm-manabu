package config

type (
	DB struct {
		User         string `envconfig:"DB_USER" default:"admin"`
		Password     string `envconfig:"DB_PASSWORD" default:"admin"`
		DatabaseName string `envconfig:"DB_DATABASE" default:"tasks"`
		Charset      string `envconfig:"DB_CHARSET" default:"utf8"`
	}
)

var (
	Database DB
)
