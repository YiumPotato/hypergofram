package controllers

import (
	"net/http"

	"github.com/YiumPotato/hypergofram/internal/db"
	"github.com/YiumPotato/hypergofram/internal/templates"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	todos, err := db.QueryMaker.ListTodos(c.Request.Context(), db.ListTodosParams{
		Limit:  100,
		Offset: 0,
	})

	if err != nil {
		c.HTML(http.StatusInternalServerError, "", templates.Error(err.Error()))
		return
	}

	c.HTML(http.StatusOK, "", templates.Home(todos))
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "", templates.About())
}
