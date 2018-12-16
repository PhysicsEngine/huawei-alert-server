package main

import (
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"github.com/PhysicsEngine/huawei-alert-server/matcher"
	"github.com/PhysicsEngine/huawei-alert-server/notification"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

type Request struct {
	macAddresses []string `json:"mac_addresses" binding:"required"`
	notification string   `json:"notification" binding:"required"`
	deviceId     string   `json:"device_id" binding:"required"`
}

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

	handler := notification.CreateHandler(logger, env)

	port := env.Port

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	matcher, err := matcher.CreateHuaweiMatcher(logger, "matcher")
	if err != nil {
		logger.Errorf("Failed to create HuaweiMatcher: %s", err)
		os.Exit(1)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/api/notification", func(c *gin.Context) {
		// TODO: Call plugin with parameter
		var req Request
		// Restore the io.ReadCloser to its original state

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Errorf("get error request:: %s, err::%s", req, err)
			// mac address can't be found
			c.JSON(400, gin.H{"status": "Invalid Request"})
			return
		}
		isMatched := false
		for _, addr := range req.macAddresses {
			logger.Infof("%s found", addr)
			if matcher.Match(addr) {
				isMatched = true
				break
			}
		}
		if isMatched {
			notify := req.notification
			if handler.Contains(notify) {
				handler.Send(notify)
				c.JSON(200, gin.H{"status": "send notification to slack"})
				return
			} else {
				logger.Errorf("not defined notification channel")
				c.JSON(400, gin.H{"status": "notfication channel not found"})
				return
			}
		}
		c.JSON(200, gin.H{"status": "target device not found"})
	})

	router.Run(":" + port)
}
