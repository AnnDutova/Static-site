package apiserver

type Config struct {
	BinAddr     string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr: ":8080",
	}
}
