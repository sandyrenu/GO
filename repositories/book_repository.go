package repositories

import "github.com/your-username/book-service/models"

var books []models.book

func GetAllBooks() []models. Book {

return books

}

func GetBookByID(id string) (models. Book, bogl){ for book ;= range books { I

if book.id ==id {

return book

}

}

return models. Book{},

}

func CreateBook (book models. Book) models. Book { book.ID = "some-unique-id" }

books #append(books, book)

return book

func UpdateBook (id string, updateBook) models. Book {

for index, book := range books {

if book.ID == id{

books [index] = upddateBook
return updateBook

}

}

return models.Book{}

}

function DeleteBook (id string) {

for index, book := range books {

if book.ID ==id {

books =append(books[:index], books[index+1:]...)
	break
}
   }
}
