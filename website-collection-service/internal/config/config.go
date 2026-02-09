package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// JWT authentication configuration
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// Database configuration
	DataSource string
}
