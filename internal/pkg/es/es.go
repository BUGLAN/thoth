package es

import (
	"github.com/olivere/elastic/v7"
)

type Config struct {
	ServerURLs []string `json:"server_urls"`
}

func Init(cfg *Config) (client *elastic.Client) {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.ServerURLs...),
		elastic.SetSniff(false), // 关闭嗅探功能
	)
	if err != nil {
		panic(err)
	}

	return client
}
