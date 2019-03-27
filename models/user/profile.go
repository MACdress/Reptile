package user

import (
	"github.com/jinzhu/gorm"
)

type Profile struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Height int `json:"height"`//身高
	Weight int `json:"weight"`//体重
	Income string  `json:"income"`//收入
	Marriage string `json:"marriage"`//婚姻状况
	Occupation string `json:"occupation"`//工作
	XinZuo string `json:"xin_zuo"`
	MyImg string `json:"my_img"`//照片
	OtherImg []string `json:"other_img"`//其他照片
	City string `gorm:"column:city"`//城市
}
type ProfileModel struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Age int `gorm:"column:age"`
	Height int `gorm:"column:height"`//身高
	Weight int `gorm:"column:weight"`//体重
	Income string  `gorm:"column:income"`//收入
	Marriage string `gorm:"column:marriage"`//婚姻状况
	Occupation string `gorm:"column:occupation"`//工作
	XinZuo string `gorm:"column:xin_zuo"`
	MyImg string `gorm:"column:my_img"`//照片
	OtherImg string`gorm:"column:other_img"`//其他照片
	City string `gorm:"column:city"`//城市
}

func (*ProfileModel)TableName()string{
	return "tb_userInfo"
}

//func CreateUserInfo(userInfo ProfileModel){
//	err := models.DBUser.Table("tb_userInfo").CreateTable(&userInfo).Error
//	if err!=nil{
//		beego.Info("插入数据失败")
//	}
//}
