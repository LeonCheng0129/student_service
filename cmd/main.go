package main

import (
	"github.com/LeonCheng0129/student_service/configs"
	"github.com/LeonCheng0129/student_service/internal/adapters/repository"
	"github.com/LeonCheng0129/student_service/internal/adapters/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// load configuration
	err := configs.NewViperConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v\n", err)
	}

	log.Printf("Initializing repository...\n")
	// 更改repo实例替换数据库具体实现
	repo := repository.NewMockRepository()

	// init server server, inject repo into handlers
	log.Printf("Starting HTTP server...\n")
	httpServer := server.NewServer(repo)

	// init gin engine
	router := gin.Default()

	// register routes
	log.Printf("Registering routes...\n")
	server.RegisterHandlers(router, httpServer)

	// start server
	addr := viper.Sub("server").GetString("addr")
	server.RunHTTPServer(addr, func(router *gin.Engine) {
		server.RegisterHandlersWithOptions(router, httpServer,
			server.GinServerOptions{
				BaseURL:      "/api",
				Middlewares:  nil,
				ErrorHandler: nil,
			})
	})
}
