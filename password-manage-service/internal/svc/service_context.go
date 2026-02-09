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
	// Initialize database connection
	db, err := gorm.Open(postgres.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database tables
	if err := db.AutoMigrate(&model.Password{}); err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		PasswordModel:  model.NewPasswordModel(db),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
