package Courses

import (
	erm "erm/Database"
)

// @Summer Course Charge
type CourseMethod struct {
	Id          int    `json:"id" gorm:"primary_key:true;unique"`
	ChargeName  string `json:"charge_name" gorm:"type:varchar(50) not null; default '';comment:'收费名称'"`
	ChargeNum   int    `json:"charge_name" gorm:"default '0';comment:'数量(课时)'"`
	TotalPrice  int    `json:"total_price" gorm:"type:decimal; default '0.00';comment:'总价(元)'"`
	ChargePrice int    `json:"charge_price" gorm:"type:decimal; default '0.00';comment:'单价(元/课时)'"`

	erm.Model
}
