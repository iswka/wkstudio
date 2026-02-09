package svc

import (
	"log"
	"user-auth-service/internal/config"
	"user-auth-service/internal/middleware"
	"user-auth-service/internal/model"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      *model.UserModel
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Initialize database connection
	db, err := gorm.Open(postgres.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database tables
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(db),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
