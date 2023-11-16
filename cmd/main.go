package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/glogger/pkg/logger"
	"github.com/yourusername/glogger/pkg/middleware"
)

var log *logger.CustomLogger

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log = logger.NewCustomLogger()


	router := gin.New()
	router.Use(middleware.RequestLoggerMiddleware(log))

	router.Run(":8080")
}
