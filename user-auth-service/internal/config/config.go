package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	
	// JWT 认证配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	
	// 数据库配置
	DataSource string
}
