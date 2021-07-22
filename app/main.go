package main

import (
	"Go-Book/app/controller"
	"Go-Book/app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	service.DBConnect()
	router := gin.Default()
	// router.Use(middleware.RecordUaAndTime)
	bookEngin := router.Group("/book")
	{
		v1 := bookEngin.Group("/api/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookEdit)
			v1.DELETE("/delete/:id", controller.BookDelete)
		}
	}
	router.Run(":8020")
}
