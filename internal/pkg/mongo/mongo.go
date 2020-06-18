package mongo

import (
	"fmt"
	"net/url"

	"github.com/globalsign/mgo"
)

type Config struct {
	// Username 用户名
	Username string `env:"DB_MONGO_USERNAME" envDefault:""`
	// Password 密码
	Password string `env:"DB_MONGO_PASSWORD" envDefault:""`
	// ServerList 服务器列表
	ServerList string `env:"DB_MONGO_SERVER_LIST" envDefault:"localhost:27017"`
	// Database 选用数据库
	Database string `env:"DB_MONGO_DATABASE" envDefault:"test"`
	// Option 其他启动选项
	Option string `env:"DB_MONGO_OPTION" envDefault:""`

	PoolLimit int `env:"DB_MONGO_POOL_LIMIT" envDefault:"50"`
}

func (config *Config) GetDsn() string {
	if config.Username != "" && config.Password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s/%s?%s",
			url.QueryEscape(config.Username),
			url.QueryEscape(config.Password),
			config.ServerList,
			config.Database,
			url.QueryEscape(config.Option))
	}
	return fmt.Sprintf("mongodb://%s/%s?%s", config.ServerList, config.Database, url.QueryEscape(config.Option))
}

func Init(config *Config) (db *mgo.Database, err error) {
	session, err := mgo.Dial(config.GetDsn())
	if err != nil {
		return nil, err
	}

	session.SetPoolLimit(config.PoolLimit)
	db = session.DB(config.Database)
	if err = db.Session.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
