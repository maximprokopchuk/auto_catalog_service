package store

type Config struct {
	DatabaseHost     string `toml:"database_host"`
	DatabaseName     string `toml:"database_name"`
	DatabaseUser     string `toml:"database_user"`
	DatabasePassword string `toml:"database_password"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseHost:     "",
		DatabaseName:     "",
		DatabaseUser:     "",
		DatabasePassword: "",
	}
}
