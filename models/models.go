package models
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
	"os"
)
var database *gorm.DB

func InitModels(){
	var err error

	//读取数据库配置信息
	cfg, err2 := ini.Load("/home/hanyuluo/go/src/yumiao/conf/server.ini")
	if err2 != nil {
		fmt.Printf("Fail to read file: %v", err2)
		os.Exit(1)
	}
	dbType:=cfg.Section("database").Key("dbType").String()
	dbName:=cfg.Section("database").Key("dbName").String()
	user:=cfg.Section("database").Key("user").String()
	password:=cfg.Section("database").Key("password").String()
	host:=cfg.Section("database").Key("host").String()


	//连接数据库
	database, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Println(err)
		panic("--- 数据库连接失败")
	}

	//InitUserModel()
	//InitBookModel()



}

func CloseDB() {
	defer database.Close()
}



func InitBookModel(){

	// Migrate the schema
	database.AutoMigrate(&Book{})


}
func InitUserModel(){

	// Migrate the schema
	database.AutoMigrate(&User{})
	database.AutoMigrate(&NoteSpace{})
	database.AutoMigrate(&Note{})

	database.Model(&NoteSpace{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	database.Model(&Note{}).AddForeignKey("note_space_id", "note_spaces(id)", "CASCADE", "CASCADE")
}

