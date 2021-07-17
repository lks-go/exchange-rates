package exchangeratesapi

func New(cfg *Config) *Client {
	return &Client{
		cfg,
	}
}

type Client struct {
	cfg *Config
}

