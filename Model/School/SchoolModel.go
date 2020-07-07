package School

import (
	erm "erm/Database"
	"fmt"
	"log"
)

// @Summer school name
type School struct {
	Id            int    `json:"id" gorm:"primary_key"`
	SchoolName    string `json:"school_name"  gorm:"type:varchar(100);not null; default ''; comment:'机构名称' " `
	SchoolAddress string `json:"school_address" gorm:"type:varchar(100) not null; default '';comment:'机构地址'"`
	ProvinceId    int    `json:"province_id" gorm:"default '0';comment:'省ID' " `
	CityId        int    `json:"city_id" gorm:"default '0';comment:'市ID' " `
	Area          int    `json:"area" gorm:"default '0';comment:'区ID' " `
	Longitude     string `json:"longitude" gorm:"type:varchar(100);not null;default '';default '';comment:'经度' " `
	Latitude      string `json:"latitude" gorm:"type:varchar(100);not null;default ''; default ''; comment:'纬度' " `
	SchoolLogo    string `json:"school_logo" gorm:"type:varchar(100);not null;default '';comment:'机构logo' "`
	SchoolDesc    string `json:"school_desc" gorm:"type:varchar(500);not null;default '';comment:'机构描述' " `
	WeChartQR     string `json:"we_chart_qr" gorm:"type:varchar(100);not null;default '';comment:'机构二维码'" `
	AdvisoryName  string `json:"advisory_name" gorm:"type:varchar(100);not null;default '';comment:'咨询名称' "`
	AdvisoryTel   string `json:"advisory_tel" gorm:"type:char(20);not null;default '';comment:'咨询电话' "`
	AdminId       int    `json:"admin_id" gorm:"default 0;comment:'机构管理员ID'"`

	erm.Model
}

// @Summer 添加机构
func AddSchool(data map[string]interface{}) (isOk bool) {
	schoolRes := erm.Db.Create(&School{
		SchoolName:    data["school_name"].(string),
		SchoolAddress: data["school_address"].(string),
		ProvinceId:    data["province_id"].(int),
		CityId:        data["city_id"].(int),
		Area:          data["area"].(int),
		Latitude:      data["latitude"].(string),
		Longitude:     data["longitude"].(string),
		SchoolLogo:    data["school_logo"].(string),
		SchoolDesc:    data["school_logo"].(string),
		WeChartQR:     data["we_chart_qr"].(string),
		AdvisoryName:  data["advisory_name"].(string),
		AdvisoryTel:   data["advisory_tel"].(string),
	})

	if schoolRes.Value != nil {
		fmt.Printf("添加学校失败:%s", schoolRes.Value)
		return false
	}
	return
}

// @Summer 编辑机构
func EditSchool(id int, data interface{}) (isOk bool) {
	editRes := erm.Db.Model(&School{}).Where("id = ?", id).Update(data)
	if editRes.Error != nil {
		log.Printf("edit dmin failed,%v", editRes)
		return false
	}
	return true
}

// @Summer 统计机构
func GetTotalSchool() (total int) {
	erm.Db.Model(&School{}).Count(&total)
	return
}

// @Summer 获取一个机构信息
func GetSchool(id int) (school School) {
	erm.Db.Related("adminInfo").Where("id = ?", id).Find(&school)
	return
}
