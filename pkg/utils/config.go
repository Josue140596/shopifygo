package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Host          string `mapstructure:"HOST"`
	Port          int    `mapstructure:"PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	Password      string `mapstructure:"PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	SSLMode       string `mapstructure:"SSLMODE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}

func GenerateDSN(c *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.DBUser, c.Password, c.DBName, c.SSLMode)
}
