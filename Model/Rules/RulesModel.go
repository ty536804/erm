package Rules

import (
	erm "erm/Database"
)

// @Summer rule
type Rule struct {
	Id       int    `json:"id" gorm:"primary_key"`
	RuleName string `json:"rule_name" gorm:"not null; default ''; comment:'角色名称' " `
	Roles    string `json:"roles" gorm:"type:text; not null; default ''; comment:'权限' "`
	IsDel    int    `json:"is_del" gorm:"default '1' ; comment:'是否删除 1有效 1软删 2禁用' " `

	erm.Model
}
