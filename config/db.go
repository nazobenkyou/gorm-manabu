package config

type (
	DB struct {
		User         string `envconfig:"DB_USER" default:"admin"`
		Password     string `envconfig:"DB_PASSWORD" default:"admin"`
		DatabaseName string `envconfig:"DB_DATABASE" default:"tasks"`
		DatabaseHost string `envconfig:"DB_HOST" default:"localhost"`
		DatabasePort string `envconfig:"DB_PORT" default:"3306"`
		Charset      string `envconfig:"DB_CHARSET" default:"utf8"`
	}
)

var (
	Database DB
)
