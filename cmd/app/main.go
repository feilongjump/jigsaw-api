package main

import (
	"fmt"
	"log"
	"strings"

	_ "jigsaw-api/docs" // for swagger
	"jigsaw-api/internal/model"
	"jigsaw-api/internal/router"
	"jigsaw-api/pkg/config"
	"jigsaw-api/pkg/database"
	"jigsaw-api/pkg/logger"
	"jigsaw-api/pkg/validator"

	"github.com/gin-gonic/gin"
)

// @title Jigsaw API
// @version 1.0
// @description Jigsaw API Server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 1. Init Config
	config.InitConfig()

	mode := strings.ToLower(config.GlobalConfig.Server.Mode)
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// 2. Init Logger
	logger.InitLogger()

	// 2.1 Init Validator
	validator.Init()

	// 3. Init Database
	database.InitDB()

	// Auto Migrate
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// 4. Init Router
	r := router.InitRouter()

	// 5. Run Server
	addr := fmt.Sprintf(":%s", config.GlobalConfig.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}
}
