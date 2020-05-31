package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Username string `form:"username" json:"username" gorm:"unique;not null"`
	Password string `form:"password" json:"password" gorm:"not null"`
	Email string `form:"email" json:"email" gorm:"unique"`
}
func CheckAuth(username, password string) bool {
	var dbuser User
	database.Where("Username = ?",username ).First(&dbuser)
	if dbuser.Password==toMd5(toMd5(password)+username){
		return true
	}
	return false
}


func GetAllUser()(users []User){
	database.Find(&users)
	return
}


func GetUser(id int)(user User){
	database.Where("id = ?", id).Find(&user)
	return
}
func GetUserByUsername(username string)(user User){
	database.Where("username = ?", username).Find(&user)
	return
}
//获取某指定用户的所有notespace
func GetUserNotespaces(id int)(notespaces []NoteSpace){
	user := GetUser(id)
	database.Where("user_id = ?", user.ID).Find(&notespaces)

	return
}

//用户新建自己的notespace
func UserAddNoteSpace(notespace NoteSpace)bool{
	err:=database.Create(&NoteSpace{
		Name: notespace.Name,
		UserID: notespace.UserID,


	}).Error
	if err==nil{
		return true
	} else{
		return false
	}
}







func Register(user User)bool{
	err:=database.Create(&User{
		Username: user.Username,
		Password: toMd5(toMd5(user.Password)+user.Username),
		Email: user.Email,

	}).Error
	if err==nil{
		return true
	} else{
		return false
	}

}

func Login(user User)bool{
	var dbuser User
	database.Where("Username = ?",user.Username ).First(&dbuser)
	if dbuser.Password==toMd5(toMd5(user.Password)+user.Username){
		return true
	}
	return false

}



func toMd5(password string) string{
	has := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", has)
}