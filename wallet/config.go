package wallet

type Config struct {
	Server struct {
		Address string `toml:"address"`
	} `toml:"server"`
	Daemon struct {
		Address  string `toml:"address"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"daemon"`
	CoinMarket struct {
		APIKEY string `toml:"api_key"`
	} `toml:"coinmarket"`
}

func DefaultConfig() Config {
	c := Config{}
	c.Daemon.Address = "127.0.0.1:8334"
	c.Server.Address = "127.0.0.1:8080"
	return c
}
