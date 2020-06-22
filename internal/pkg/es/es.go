package es

import (
	"github.com/olivere/elastic/v7"
)

type Config struct {
	ServerURLs []string `json:"server_urls"`
}

func Init(cfg *Config) (client *elastic.Client, err error) {
	client, err = elastic.NewClient(
		elastic.SetURL(cfg.ServerURLs...),
		elastic.SetSniff(false), // 关闭嗅探功能
	)
	if err != nil {
		return nil, err
	}
	// todo auth
	return client, err
}
