package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"go.uber.org/zap"
)

func main() {
	// setup logger
	zapLogger, _ := zap.NewProduction()

	logger := zapLogger.Sugar()
	defer func() {
		err := zapLogger.Sync()
		log.Fatal(err)
	}()

	env, err := config.ReadFromEnv()
	if err != nil {
		logger.Errorf("Failed to read env vars: %s", err)
		os.Exit(1)
	}

	port := env.Port

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
