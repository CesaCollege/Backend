package config

type MysqlDB struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}
type Config struct {
	DB MysqlDB
}
