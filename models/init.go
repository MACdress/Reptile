package models

import (
	"Reptile/models/user"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DBUser * gorm.DB
func init (){
	DBUser = MasterDbConnection()
	DBUser.CreateTable(&user.ProfileModel{})
}



func MasterDbConnection() *gorm.DB{
	db_driver := beego.AppConfig.String("db_driver")
	db_user := beego.AppConfig.String("db_user")
	db_passwd := beego.AppConfig.String("db_password")
	db_host := beego.AppConfig.String("db_host")
	db_port, _ := beego.AppConfig.Int("db_port")
	db_database_name := beego.AppConfig.String("db_database_name")
	MasterDB, err := gorm.Open(db_driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",db_user,db_passwd,db_host,db_port,db_database_name))
	if err!=nil{
		logs.Info("Master写数据库连接失败")
		logs.Info(err)
	}
	return MasterDB
}
