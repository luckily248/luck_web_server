package luckweb

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luckily248/luck_web_server/handler"
)

func Run() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./static/*.html")
	router.GET("/", handler.Homeh)
	router.Run(":" + os.Getenv("PORT"))
}
