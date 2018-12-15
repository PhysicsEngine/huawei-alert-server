package main

import (
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

type Request struct {
	mac_addresses []string
	notification  string
	device_id     string
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

	port := env.Port

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	matcher, err := matcher.CreateHuaweiMatcher(logger)
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
		if err := c.ShouldBindJSON(&req); err != nil {
			// mac address can't be found
			c.JSON(400, gin.H{"status": "Invalid Request"})
		}
		is_huawei_detected := false
		for addr := range req.mac_addresses {
			if matcher.Match(addr) {
				is_huawei_detected = true
				break
			}
		}
		if is_huawei_detected {
			notifyDevice := "slack"
			jsonStr := "{}"
			switch notifyDevice {
			case
				"slack":
				slackhandler.PostSlack(jsonStr, logger)
			default:
				logger.Errorf("no device notified")
			}
		}
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run(":" + port)
}
