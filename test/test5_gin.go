package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main5() {
	router := gin.Default()
	router.GET("/time", func(c *gin.Context) {
		m := map[string]string{
			"time": time.Now().Format("2006-01-02"),
		}
		c.JSON(http.StatusOK, m)
	})
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
