package main

import (
	"github/teacoder-team/golang-backend/config"
	"github/teacoder-team/golang-backend/routes"
	"log"
	"strconv"

	_ "github/teacoder-team/golang-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Response struct {
	Message string `json:"message"`
}

// @title TeaCoder API
// @description API for Teacoder educational platform
// @version 1.0.0

// @contact.name TeaCoder Team
// @contact.url https://teacoder.ru
// @contact.email help@teacoder.ru
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	config.ConnectDatabase(cfg)
	config.ConnectRedis(cfg)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.RegisterAccountRoutes(router)

	router.SetTrustedProxies(nil)

	if enabled := cfg.SwaggerEnabled; enabled {
		router.GET(cfg.SwaggerPrefix+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Printf("📄 Swagger documentation is available at: %s%s/index.html\n", cfg.ApplicationURL, cfg.SwaggerPrefix)
	}

	log.Printf("🚀 Server is running at: %s\n", cfg.ApplicationURL)

	if err := router.Run(":" + strconv.Itoa(cfg.ApplicationPort)); err != nil {
		log.Fatalf("❌ Error starting server: %v\n", err)
	}
}
