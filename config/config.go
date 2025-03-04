package config

import (
  "github.com/spf13/viper"
)

type Config struct {
  AppPort int
  DbPort int
  DbHost string
  DbPassword string
  DbName string
  DbUser string
}

func setDefaults() {
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("config")
  viper.AutomaticEnv()

  viper.SetDefault("app.port", 8000)
  viper.SetDefault("db.port", 5432)
  viper.BindEnv("db.password", "DB_PASSWORD")
  viper.SetDefault("db.host", "db")
  viper.SetDefault("db.name", "postgres")
  viper.SetDefault("db.user", "postgres")
}

func Default() *Config {
  setDefaults()

  viper.ReadInConfig()

  appPort := viper.GetInt("app.port")
	dbPort := viper.GetInt("db.port")
	dbHost := viper.GetString("db.host")
	dbUser := viper.GetString("db.user")
  dbName := viper.GetString("db.name")
  dbPassword := viper.GetString("db.password")

  return &Config {
    AppPort: appPort,
    DbPort: dbPort,
    DbUser: dbUser,
    DbName: dbName,
    DbHost: dbHost,
    DbPassword: dbPassword, 
  }
}
