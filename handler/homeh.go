package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homeh(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "luke's websit",
	})
}
