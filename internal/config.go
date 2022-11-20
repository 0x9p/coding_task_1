package internal

import "os"

type DbConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

type Config struct {
	ServiceName string
	MasterDb    DbConfig
}

func NewConfig() *Config {
	return &Config{
		ServiceName: LookupEnv("SERVICE_NAME", "YOUR_SERVICE_NAME"),
		MasterDb: DbConfig{
			Host:     LookupEnv("MASTER_DATABASE_HOST", "YOUR_DATABASE_HOST"),
			Name:     LookupEnv("MASTER_DATABASE_NAME", "YOUR_DATABASE_NAME"),
			User:     LookupEnv("MASTER_DATABASE_USER", "YOUR_DATABASE_USER"),
			Password: LookupEnv("MASTER_DATABASE_PASSWORD", "YOUR_DATABASE_PASSWORD"),
		},
	}
}

func LookupEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
