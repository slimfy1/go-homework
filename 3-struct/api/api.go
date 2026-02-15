package api

import "homework/3-struct/config"

type Client struct {
	config  *config.Config
	baseURL string // например "https://api.jsonbin.io/v3"
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		config:  cfg,
		baseURL: "https://api.jsonbin.io/v3",
	}
}

func (ac *Client) CreateBin(data []byte) error {
	// Здесь используй ac.config.Key для заголовка X-Master-Key
	// Отправь POST запрос на ac.baseURL + "/b"
	return nil
}

func (ac *Client) GetBin(binId string) ([]byte, error) {
	// GET запрос с ключом из config
	return nil, nil
}
