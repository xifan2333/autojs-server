package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, AutoJSPro!",
		})
	})

	r.GET("/api/v1/account", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":           uuid.New().String(),
			"now":          time.Now().UnixMilli(),
			"emailAddress": "Auto@xifan.fun",
			"fullName":     "AutoJSPro",
			"paidServices": gin.H{
				"v8": gin.H{
					"expires": time.Now().Add(365 * 24 * time.Hour).UnixMilli(),
				},
			},
			"permissions": gin.H{},
		})
	})

	// 使用环境变量配置监听地址，默认为 0.0.0.0:8080
	listenAddr := "0.0.0.0:8080"

	r.Run(listenAddr)
}
