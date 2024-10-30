package config

import "os"

type Config struct {
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBPORT     string
}

func LoadConfig() *Config {
	return &Config{
		DBPORT:     getEnv("DB_PORT", "3306"),
		DBUSER:     getEnv("DB_USER", "root"),
		DBPASSWORD: getEnv("DB_PASSWORD", "root"),
		DBHOST:     getEnv("DB_HOST", "localhost"),
		DBNAME:     getEnv("DB_NAME", "urlShortener"),
	}
}

func getEnv(key, defaulVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	return defaulVal
}
