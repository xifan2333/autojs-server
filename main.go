package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/account", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":           "6131f76468e4553fba39ae4c",
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

	r.Run(":80")
}
