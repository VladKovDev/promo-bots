package config

func SetDefaultConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:         "localhost",
			Port:         5432,
			User:         "postgres",
			Password:     "12345",
			Name:         "promo-bots",
			SSLMode:      "disable",
			MaxOpenConns: 25,
			MaxIdleConns: 500,
		},
		Logger: LoggerConfig{
			Level:  "debug",
			Pretty: true,
		},
	}
}
