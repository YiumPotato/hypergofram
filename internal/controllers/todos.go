package controllers

import (
	"fmt"
	"strconv"

	"github.com/YiumPotato/hypergofram/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func errorResponse(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
}

/* Classic JSON API routes, for other services to consume
v1.GET("todo", Controllers.GetTodos)
v1.POST("todo", Controllers.CreateATodo)
v1.GET("todo/:id", Controllers.GetATodo)
v1.PUT("todo/:id", Controllers.UpdateATodo)
v1.DELETE("todo/:id", Controllers.DeleteATodo)
*/

// Note: This file consists of simple CRUD operations for the Todo model, returning JSON
// it will not be used in the hypergofram web app, but is here for reference
// and to show how to use the db package

func GetTodos(c *gin.Context) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Println("Error with Limit Query")
		errorResponse(c, err)
		return
	}

	offsetStr := c.Query("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		fmt.Println("Error with Limit Query")
		errorResponse(c, err)
		return
	}

	todos, err := db.QueryMaker.ListTodos(c, db.ListTodosParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		fmt.Println("Error with List Query")
		errorResponse(c, err)
		return
	}

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func CreateATodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		fmt.Println("Error with Create Query")
		errorResponse(c, fmt.Errorf("title cannot be empty"))
		return
	}

	todo, err := db.QueryMaker.CreateTodo(c, db.CreateTodoParams{
		Title: pgtype.Text{String: title, Valid: true},
		Done:  pgtype.Bool{Bool: false, Valid: true},
	})

	if err != nil {
		fmt.Println("Error with Create Query")
		errorResponse(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"todo": todo,
		"err":  err,
	})
	// c.Redirect(302, "/")

}

func GetATodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing todo ID")
		errorResponse(c, err)
		return
	}

	todo, err := db.QueryMaker.GetTodo(c, todoID)

	if err != nil {
		fmt.Println("Error with Get Query")
		errorResponse(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func UpdateATodo(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	done := c.PostForm("done")

	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing todo ID")
		errorResponse(c, err)
		return
	}

	todo, err := db.QueryMaker.UpdateTodo(c, db.UpdateTodoParams{
		ID:    todoID,
		Title: pgtype.Text{String: title, Valid: true},
		Done:  pgtype.Bool{Bool: done == "true", Valid: true},
	})

	if err != nil {
		fmt.Println("Error with Update Query")
		errorResponse(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func DeleteATodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error parsing todo ID")
		errorResponse(c, err)
		return
	}

	err = db.QueryMaker.DeleteTodo(c, todoID)

	if err != nil {
		fmt.Println("Error with Delete Query")
		errorResponse(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
