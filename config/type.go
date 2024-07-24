package config

type (
	ApiConfig struct {
		Url string
	}

	DbConfig struct {
		DataSourceName string
	}

	Config struct {
		ApiConfig
		DbConfig
	}
)
