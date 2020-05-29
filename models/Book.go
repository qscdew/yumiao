package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type Book struct{
	gorm.Model

	Title string `form:"title" json:"title" gorm:"type:varchar(50)"`
	Author string `form:"author" json:"author"gorm:"type:varchar(50)"`
	Publisher string `form:"publisher" json:"publisher" gorm:"type:varchar(50)"`
	OriginalTitle string `form:"otitle" json:"otitle" gorm:"type:varchar(50)"`
	Translator string `form:"translator" json:"translator" gorm:"type:varchar(50)"`
	PublicationDate *time.Time `form:"date" json:"date"`
	Pricing sql.NullFloat64 `form:"pricing" json:"pricing"`
	Binding string `form:"binding" json:"binding" gorm:"type:varchar(10)"`
	Series string `form:"series" json:"series" gorm:"type:varchar(50)"`
	ISBN string `form:"isbn" json:"isbn"`
	Brief string `form:"brief" json:"brief"`
	Image string `form:"image" json:"image"`
	DoubanRating sql.NullFloat64
	DoubanRatingNum int

}

func GetBook(id int)(book Book){
	database.Where("id = ?", id).First(&book)
	return
}

func GetAllBook()(books []Book){
	database.Find(&books)
	return
}
func AddBook(book Book)bool{
	err:=database.Create(&Book{
		Title:book.Title,
		Author: book.Author,
		Publisher :book.Publisher,
		OriginalTitle :book.OriginalTitle,
		Translator:book.Translator,
		PublicationDate :book.PublicationDate,
		Pricing :book.Pricing,
		Binding:book.Binding,
		Series :book.Series,
		ISBN :book.ISBN,
		Brief:book.Brief,
		Image :book.Image,

	}).Error
	if err==nil{
		return true
	} else{
		return false
	}
}

func DeleteBook(id int)bool{
	var book Book
	database.Where("id = ?", id).Delete(&book)
	return true
}

func EditBook(id int,book Book)bool{
	database.Model(&book).Where("id = ?", id).Updates(book)
	return true
}