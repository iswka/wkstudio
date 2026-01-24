package svc

import (
	"log"
	"user-auth-service/internal/config"
	"user-auth-service/internal/middleware"
	"user-auth-service/internal/model"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      *model.UserModel
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	// 自动迁移数据表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("数据表迁移失败: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(db),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
