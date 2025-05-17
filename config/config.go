package config

// AppConfig holds the application configuration
type AppConfig struct {
	Port        string
	Environment string
	Version     string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
}

// GetConfig returns the application configuration
func GetConfig() *AppConfig {
	return &AppConfig{
		Port:        "8080",
		Environment: "development",
		Version:     "1.0.0",
		DBHost:      "localhost",
		DBPort:      "5432",
		DBUser:      "cmsuser",
		DBPassword:  "",
		DBName:      "cmsdb",
		JWTSecret:   "",
	}
}