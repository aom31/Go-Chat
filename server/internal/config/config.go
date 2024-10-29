package config

import "github.com/spf13/viper"

type Container struct {
	DBPostgres *DBPostgres
}

type DBPostgres struct {
	User     string
	Password string
	Port     int
	Name     string
	Host     string
}

func NewConfig() (*Container, error) {
	viper.AutomaticEnv()

	dbPostgres := DBPostgres{
		User:     viper.GetString("PostgresUser"),
		Password: viper.GetString("PostgresPassword"),
		Port:     viper.GetInt("PostgresPort"),
		Name:     viper.GetString("DBName"),
		Host:     viper.GetString("DBHost"),
	}

	return &Container{
		DBPostgres: &dbPostgres,
	}, nil
}
