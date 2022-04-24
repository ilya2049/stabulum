package config

type Config struct {
	API      APIConfig
	Postgres PostgresConfig
}

type APIConfig struct {
	HTTPServer APIHTTPServerConfig
}

type APIHTTPServerConfig struct {
	Address string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
