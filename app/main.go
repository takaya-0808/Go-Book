package main

import (
	"Go-Book/app/controller"
	"Go-Book/app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := service.DBConnect()
	router := gin.Default()
	// router.Use(middleware.RecordUaAndTime)
	bookController := controller.NewBookController(*db)
	bookEngin := router.Group("/book")
	{
		v1 := bookEngin.Group("/api/v1")
		{
			v1.POST("/add", bookController.BookAdd)
			v1.GET("/list", bookController.BookList)
			v1.PUT("/update", bookController.BookEdit)
			v1.DELETE("/delete/:id", bookController.BookDelete)
			v1.GET("/list/:id", bookController.BookSearch)
		}
	}
	router.Run(":8020")
}
