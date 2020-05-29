package models

import "github.com/jinzhu/gorm"

type NoteSpace struct {
	gorm.Model

	Name string
	UserID uint
}
//获取某指定notespace
func GetNoteSpace(id int)(NoteSpace NoteSpace){
	database.Where("id = ?", id).First(&NoteSpace)
	return
}
//获取所有notespace
func GetAllNoteSpaces()(NoteSpaces []NoteSpace){
	database.Find(&NoteSpaces)
	return
}

//获取某指定的notespace下的所有note
func GetNoteSpaceNotes(id int)(Notes []Note){

	database.Where("note_space_id = ?", id).Find(&Notes)
	return
}

//给某指定的notespace增加note
func AddNote(note Note)bool{
	err:=database.Create(&Note{
		Text: note.Text,
		NoteSpaceID: note.NoteSpaceID,

	}).Error
	if err==nil{
		return true
	} else{
		return false
	}
}

