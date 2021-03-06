package Admin

import (
	crm "erm/Database"
	"erm/Model/School"
	"erm/Pkg/Setting"
	"log"
)

type SysAdminUser struct {
	Id           int             `json:"id" gorm:"primary_key:true;unique"`
	LoginName    string          `json:"login_name" grom:"type:varchar(100);not null; default ''; comment:'账号'" `
	Pwd          string          `json:"pwd" gorm:"varchar(32); not null; default '';comment:'密码'" `
	Status       int             `json:"status" gorm:"not null;default 1;comment:'状态 1 有效 0 无效'"`
	CityID       int             `json:"city_id" gorm:"not null;default 0; comment:'城市'"`
	Avatar       string          `json:"avatar" gorm:"type:varchar(100);not null; default ''; comment:'头像'" `
	Tel          string          `json:"tel" gorm:"type:varchar(20); not null; default '';comment:'电话' "`
	Email        string          `json:"email" gorm:"type:varchar(200); not null; default ''; comment:'邮箱'"`
	NickName     string          `json:"nick_name" gorm:"type:varchar(100); not null; default ''; comment:'昵称' "`
	DepartmentId string          `json:"department_id" gorm:"type:varchar(200);not null; default ''; comment:'部门ID'" `
	PowerId      string          `json:"power_id" gorm:"type:text; not null; default ''; comment:'权限'" `
	School       []School.School `gorm:"ForeignKey:AdminId;AssociationForeignKey:Id"`

	crm.Model
}

// @Summer 添加管理员
func AddAdminUser(data map[string]interface{}) (isOk bool) {
	adminRes := crm.Db.Create(&SysAdminUser{
		LoginName: data["login_name"].(string),
		Pwd:       data["login_name"].(string),
		Status:    data["status"].(int),
		CityID:    data["city_id"].(int),
		Avatar:    data["avatar"].(string),
		Tel:       data["tel"].(string),
		Email:     data["email"].(string),
		NickName:  data["nick_name"].(string),
	})

	if adminRes.Error != nil {
		log.Printf("add dmin failed,%v", adminRes)
		return false
	}
	return
}

// @Summer 编辑管理员
func EditAdminUser(id int, data interface{}) (isOk bool) {
	editRes := crm.Db.Model(&SysAdminUser{}).Where("id = ?", id).Update(data)
	if editRes.Error != nil {
		log.Printf("edit dmin failed,%v", editRes)
		return false
	}
	return
}

// @Summer 获取单条管理员信息
func GetAdminUser(id int) (admin SysAdminUser) {
	crm.Db.Preload("School").Where("id = ?", id).Find(&admin)
	return
}

// @Summer 统计管理员信息
func GetTotalAdmin() (count int) {
	crm.Db.Where(&SysAdminUser{}).Count(&count)
	return
}

//获取管理员列表
func GetAdminUserList(page int, where interface{}) (admin []SysAdminUser) {
	offset := 0
	if page >= 1 {
		offset = (page - 1) * Setting.PageSize
	}
	crm.Db.Where(where).Limit(Setting.PageSize).Offset(offset).Find(&admin)
	return
}

// @Summer 判断当前账号是否已经注册
func ExistsByLoginName(loginName string) (isOk bool) {
	var user SysAdminUser
	crm.Db.Select("id").Where("login_name = ?", loginName).First(&user)
	if user.Id < 1 {
		return false
	}
	return
}

// @Summer 通过用户名获取当前用户信息
func GetAdmin(uname string) (admin SysAdminUser) {
	crm.Db.Where("login_name = ?", uname).First(&admin)
	return
}
