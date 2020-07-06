package Admin

import (
	erm "erm/Database"
	"erm/Pkg/Setting"
	"log"
)

type SysAdminUser struct {
	Id           int    `json:"id" gorm:"primary_key"`
	LoginName    string `json:"login_name" grom:"type:varchar(100);not null; default ''; comment:'账号'" `
	Pwd          string `json:"pwd" gorm:"varchar(32); not null; default '';comment:'密码'" `
	Status       int    `json:"status" gorm:"not null;default 1;comment:'状态 1 有效 0 无效'"`
	CityID       int    `json:"city_id" gorm:"default '0'; comment:'城市'"`
	Avatar       string `json:"avatar" gorm:"type:varchar(100);not null; default ''; comment:'头像'" `
	Tel          string `json:"tel" gorm:"type:varchar(20); not null; default '';comment:''电话 "`
	Email        string `json:"email" gorm:"type:varchar(200); not null; default ''; comment:' 邮箱'"`
	NickName     string `json:"nick_name" gorm:"type:varchar(100); not null; default ''; comment: '昵称' "`
	DepartmentId int    `json:"department_id" gorm:"default '0'; comment:'部门ID'" `
	PowerId      string `json:"power_id" gorm:"type:text; not null; default ''; comment:'权限'" `
	SchoolName   string `json:"school_name" gorm:"type:varchar(100);not null; default ''; comment:'学校名称' " `

	erm.Model
}

// @Summer 添加管理员
func AddAdminUser(data map[string]interface{}) (isOk bool) {
	adminRes := erm.Db.Create(&SysAdminUser{
		LoginName:  data["login_name"].(string),
		Pwd:        data["login_name"].(string),
		Status:     data["status"].(int),
		CityID:     data["city_id"].(int),
		Avatar:     data["avatar"].(string),
		Tel:        data["tel"].(string),
		Email:      data["email"].(string),
		NickName:   data["nick_name"].(string),
		SchoolName: data["login_name"].(string),
	})

	if adminRes.Error != nil {
		log.Printf("add dmin failed,%v", adminRes)
		return false
	}
	return true
}

// @Summer 编辑管理员
func EditAdminUser(id int, data interface{}) (isOk bool) {
	editRes := erm.Db.Model(&SysAdminUser{}).Where("id = ?", id).Update(data)
	if editRes.Error != nil {
		log.Printf("edit dmin failed,%v", editRes)
		return false
	}
	return true
}

// @Summer 获取单条管理员信息
func GetAdminUser(id int) (admin SysAdminUser) {
	erm.Db.Where("id = ?", id).Find(&admin)
	return
}

// @Summer 统计管理员信息
func GetTotalAdmin() (count int) {
	erm.Db.Where(&SysAdminUser{}).Count(&count)
	return
}

//获取管理员列表
func GetAdminUserList(page int, where interface{}) (admin []SysAdminUser) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * Setting.PageSize
	}
	erm.Db.Where(where).Limit(Setting.PageSize).Offset(offset).Find(&admin)
	return
}
