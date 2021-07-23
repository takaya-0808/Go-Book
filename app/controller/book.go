package controller

import (
	"Go-Book/app/model"
	"Go-Book/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(sql service.SqlHandler) *BookController {
	return &BookController{
		bookService: service.BookService{
			SqlHandler: sql,
		},
	}
}

func (bc *BookController) BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	id, err := bc.bookService.SetBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"ID": id})
}

func (bc *BookController) BookList(c *gin.Context) {

	BookLists := bc.bookService.GetBookList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

func (bc *BookController) BookEdit(c *gin.Context) {
	book := model.Book{}
	c.Bind(&book)
	err := bc.bookService.UpdateBook(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "ok"})
}

func (bc *BookController) BookDelete(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err = bc.bookService.DeleteBook(int(intID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "ok"})
}
