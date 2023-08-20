package config

type Config struct {
	Database struct {
		Host     string `env:"DATABASE_HOST" env-default:"localhost"`
		Port     int    `env:"DATABASE_PORT" env-default:"5432"`
		Name     string `env:"DATABASE_NAME" env-default:"book_manager_db"`
		Username string `env:"DATABASE_USERNAME" env-default:"admin"`
		Password string `env:"DATABASE_PASSWORD" env-default:"admin"`
	}
	Redis struct {
		Host string `env:"REDIS_HOST" env-default:"localhost"`
		Port int    `env:"REDIS_PORT" env-default:"6379"`
	}
}
