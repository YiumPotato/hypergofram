package controllers

import (
	"net/http"

	"github.com/YiumPotato/hypergofram/internal/templates"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "", templates.Home())
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "", templates.About())
}
