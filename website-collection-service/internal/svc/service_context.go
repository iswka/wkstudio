package svc

import (
	"log"
	"website-collection-service/internal/config"
	"website-collection-service/internal/middleware"
	"website-collection-service/internal/model"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	WebsiteModel   *model.WebsiteModel
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Initialize database connection
	db, err := gorm.Open(postgres.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database tables
	if err := db.AutoMigrate(&model.Website{}); err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}

	return &ServiceContext{
		Config:         c,
		WebsiteModel:   model.NewWebsiteModel(db),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
