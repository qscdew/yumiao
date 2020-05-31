package models

import "github.com/jinzhu/gorm"

type Note struct{
	gorm.Model

	NoteSpaceID uint
	UserID uint
	Text string
}
//获取某指定的note
func GetNote(id int)(note Note){
	database.Where("id = ?", id).First(&note)
	return
}

//获取所有note
func GetAllNotes()(notes []Note){
	database.Find(&notes)
	return
}
//删除某条指定的note
func DeleteNote(id int)bool{
	var note Note
	database.Where("id = ?", id).Delete(&note)

	return true
}

//修改某条指定的note
func EditNote(id int,note Note)bool{
	database.Model(&note).Where("id = ?", id).Updates(note)
	return true
}
