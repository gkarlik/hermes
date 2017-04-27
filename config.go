package hermes

type Config struct {
	Username       string
	ServerEndpoint string
}

func NewConfig(username, serverEndpoint string) *Config {
	return &Config{
		Username:       username,
		ServerEndpoint: serverEndpoint,
	}
}
