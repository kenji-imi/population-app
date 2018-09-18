package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kenji-imi/population-app/src/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	// DB Open
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3309)/populationdb")
	if err != nil {
		panic(err)
	}

	h := handler.Handler{
		DB: db,
	}

	r := gin.Default() // router
	r.Use(sampleMiddleware())
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/population", h.Create)
	r.PATCH("/population/:id", h.Update)
	r.DELETE("/population/:id", h.Delete)
	r.GET("/population/:id", h.Read)
	r.GET("/population", h.List)

	r.Run("127.0.0.1:10080")
}

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
