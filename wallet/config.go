package wallet

type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`
	Daemon struct {
		Address  string `json:"address"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"daemon"`
	CoinMarket struct {
		APIKEY string `json:"api_key"`
	} `json:"coinmarket"`
}

func DefaultConfig() Config {
	c := Config{}
	c.Daemon.Address = "127.0.0.1:8334"
	return c
}
