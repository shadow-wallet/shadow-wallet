package wallet

type Config struct {
	Daemon struct {
		Address  string
		Username string
		Password string
	}
}

func DefaultConfig() Config {
	c := Config{}
	c.Daemon.Address = "127.0.0.1:8334"
	c.Daemon.Username = "user"
	c.Daemon.Password = "pass"
	return c
}
