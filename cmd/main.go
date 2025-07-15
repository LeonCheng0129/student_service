package main

import (
	"github.com/LeonCheng0129/student_service/internal/adapters/repository"
	"github.com/LeonCheng0129/student_service/internal/adapters/server"
	_ "github.com/LeonCheng0129/student_service/internal/common/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Printf("Initializing repository...\n")
	// 更改repo实例替换数据库具体实现
	//repo := repository.NewMockRepository()
	repo, err := repository.NewMySQLRepository()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v\n", err)
	}
	cachedRepo := repository.NewCachedRepository(repo)

	// init server server, inject repo into handlers
	log.Printf("Starting HTTP server...\n")
	httpServer := server.NewServer(cachedRepo)

	// init gin engine
	router := gin.Default()

	// register routes
	log.Printf("Registering routes...\n")
	server.RegisterHandlersWithOptions(router, httpServer, server.GinServerOptions{
		BaseURL:      "/api",
		Middlewares:  nil,
		ErrorHandler: nil,
	})

	// run server
	addr := viper.GetString("http.addr")
	if addr == "" {
		log.Fatal("HTTP server address is not configured")
	}
	log.Printf("Running HTTP server on %s...\n", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run HTTP server: %v\n", err)
	}
}
