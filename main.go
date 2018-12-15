package main

import (
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/zap"
	"log"
	"net/http"
	"bytes"	
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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/api/notification", func(c *gin.Context) {
		// TODO: Call plugin with parameter
		jsonStr := "{}"
		url := "https://maker.ifttt.com/trigger/huawei_alert/with/key/c9GxSBX5gGyKITjQTGsuwH"
    req, err := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer([]byte(jsonStr)),
		)		
		if err != nil {
			logger.Errorf("Invalid http request")
		}
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Errorf("Fail to send notification")
		}		
    defer resp.Body.Close()		

		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run(":" + port)
}
