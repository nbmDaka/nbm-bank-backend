package config

type Config struct {
	App AppConfig

	Keycloak KeycloakConfig
}

type AppConfig struct {
	Port string
}

type KeycloakConfig struct {
	URL string

	Realm string

	ClientID string

	ClientSecret string
}

func Load() Config {

	return Config{

		App: AppConfig{
			Port: getEnv(
				"APP_PORT",
				"8081",
			),
		},

		Keycloak: KeycloakConfig{

			URL: getEnv(
				"KEYCLOAK_URL",
				"",
			),

			Realm: getEnv(
				"KEYCLOAK_REALM",
				"",
			),

			ClientID: getEnv(
				"KEYCLOAK_CLIENT_ID",
				"",
			),

			ClientSecret: getEnv(
				"KEYCLOAK_CLIENT_SECRET",
				"",
			),
		},
	}
}
