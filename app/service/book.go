package service

import (
	"Go-Book/app/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BookService struct {
	SqlHandler
}

func (bs *BookService) SetBook(book *model.Book) (id int, err error) {

	result, err := bs.Conn.Exec("Insert into books (title, content) values (?, ?)", book.Title, book.Content)

	if err != nil {
		panic(err.Error())
	}

	ID64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(ID64)
	return
}

func (bs *BookService) GetBookList() []model.Book {

	rows, err := bs.Conn.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		book := model.Book{}
		if err := rows.Scan(&book.Id, &book.Title, &book.Content); err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func (bs *BookService) UpdateBook(book *model.Book) error {

	bookUpdate, err := bs.Conn.Exec("update books set title=? and content=? where id=?", &book.Title, &book.Content, &book.Id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := bookUpdate.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	log.Println(rowsAffect)
	return err
}

func (bs *BookService) DeleteBook(id int) error {
	bookDelete, err := bs.Conn.Exec("delete from books where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := bookDelete.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	log.Println(rowsAffect)
	return err
}
