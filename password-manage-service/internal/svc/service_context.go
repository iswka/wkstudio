package svc

import (
	"log"
	"password-manage-service/internal/config"
	"password-manage-service/internal/middleware"
	"password-manage-service/internal/model"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	PasswordModel  *model.PasswordModel
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := gorm.Open(postgres.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	// 自动迁移数据表
	if err := db.AutoMigrate(&model.Password{}); err != nil {
		log.Fatalf("数据表迁移失败: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		PasswordModel:  model.NewPasswordModel(db),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
