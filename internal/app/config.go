package api

// General instance for API server of REST application
type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	 
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
