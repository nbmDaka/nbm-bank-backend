package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Environment string
	Port        string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() Config {

	return Config{

		App: AppConfig{

			Environment: GetEnv(
				"APP_ENV",
				"development",
			),

			Port: GetEnv(
				"APP_PORT",
				"50051",
			),
		},

		Database: DatabaseConfig{

			Host: GetEnv(
				"POSTGRES_HOST",
				"localhost",
			),

			Port: GetEnv(
				"POSTGRES_PORT",
				"5432",
			),

			User: GetEnv(
				"POSTGRES_USER",
				"postgres",
			),

			Password: GetEnv(
				"POSTGRES_PASSWORD",
				"",
			),

			Name: GetEnv(
				"POSTGRES_DB",
				"postgres",
			),
		},
	}
}
