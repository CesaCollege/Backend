package config

type PostgresDB struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}
type Config struct {
	DB PostgresDB
}
