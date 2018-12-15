package main

import (
	"bytes"
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"github.com/PhysicsEngine/huawei-alert-server/matcher"
	"github.com/PhysicsEngine/huawei-alert-server/slackhandler"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
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

	matcher, err := createHuaweiMatcher()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/api/notification", func(c *gin.Context) {
		// TODO: Call plugin with parameter
		is_huawei_detected := false
		mac_addresses, err := c.Request.Body.get("mac_addresses")
		for _, addr := range mac_addresses {
			if matcher.match(addr) {
				is_huawei_detected = true
				break
			}
		}
		if is_huawei_detected {
			notifyDevice := "slack"
			switch notifyDevice {
				case "slack": slackhandler.PostSlack(jsonStr)
				defaut: logger.error("no device notified") 
			}
			jsonStr := "{}"
			slackhandler.PostSlack(jsonStr)
		}
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run(":" + port)
}
