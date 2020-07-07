package Sales

import (
	erm "erm/Database"
	"erm/Pkg/Setting"
	"fmt"
)

// @Summer 线索管理
type Student struct {
	Id             int    `json:"id" gorm:"primary_key"`
	Avatar         string `json:"avatar" gorm:"type:varchar(100); not null ; default ''; comment:'头像' " `
	StudentName    string `json:"student_name" gorm:"type:varchar(100); not null ; default ''; comment:'学生姓名' " `
	ContactName    string `json:"contact_name" gorm:"type:varchar(100); not null ; default ''; comment:'联系人' " `
	ContactTel     string `json:"contact_tel" gorm:"type:varchar(20); not null ; default ''; comment:'联系人电话' " `
	StudentSex     int    `json:"student_sex" gorm:"default 0; comment:'学生性别 0未知 1男 2女' "`
	StudentAge     int    `json:"student_age" gorm:"default 0; comment:'学生年龄' " `
	Birthday       string `json:"birthday" gorm:"type:char(10)not null;default ''; comment:'生日' " `
	CurrentSchool  string `json:"current_school" gorm:"type:varchar(100); not null ; default ''; comment:'当前学校' "`
	CurrentClass   string `json:"current_class" gorm:"type:varchar(100); not null ; default ''; comment:'当前年级' "`
	SpareName      string `json:"spare_name" gorm:"type:varchar(100); not null ; default ''; comment:'第二联系人' "`
	SpareTel       string `json:"spare_tel" gorm:"type:varchar(100); not null ; default ''; comment:'第二联系人电话' "`
	IntentionLevel int    `json:"intention_level" gorm:"default '0'; comment:'意向级别' "`
	StudentCom     string `json:"student_com" gorm:"type:varchar(100); not null ; default ''; comment:'学员来源' " `
	Follower       string `json:"follower" gorm:"type:varchar(100); not null ; default ''; comment:'跟进人' "`
	Remark         string `json:"remark" gorm:"type:varchar(500); not null ; default ''; comment:'备注' " `
	Span           string `json:"span" gorm:"type:varchar(200); not null ; default ''; comment:'标签' "`

	erm.Model
}

// @Summer 添加线索
func AddStudent(data map[string]interface{}) (isOk bool) {
	stuRes := erm.Db.Create(&Student{
		Avatar:         data["avatar"].(string),
		StudentName:    data["student_name"].(string),
		ContactName:    data["contact_name"].(string),
		ContactTel:     data["contact_tel"].(string),
		StudentSex:     data["student_sex"].(int),
		StudentAge:     data["student_age"].(int),
		Birthday:       data["birthday"].(string),
		CurrentSchool:  data["current_school"].(string),
		CurrentClass:   data["current_class"].(string),
		SpareName:      data["spare_name"].(string),
		SpareTel:       data["spare_tel"].(string),
		IntentionLevel: data["intention_level"].(int),
		StudentCom:     data["student_com"].(string),
		Follower:       data["follower"].(string),
		Remark:         data["remark"].(string),
		Span:           data["span"].(string),
	})
	if stuRes.Error != nil {
		fmt.Printf("add fiald:%s", stuRes.Error)
		return false
	}
	return
}

// 编辑线索
func EditStudent(id int) (stu Student) {
	erm.Db.Where("id = ?", id).Find(&stu)
	return
}

// @Summer 获取线索列表
func GetStudent(page int, where interface{}) (stu []Student) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * Setting.PageSize
	}
	erm.Db.Model(where).Limit(Setting.PageSize).Offset(offset).Find(&stu)
	return
}
