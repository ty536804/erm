package Sales

import (
	erm "erm/Database"
	"time"
)

// @Summer 跟进记录
type Follower struct {
	Id           int       `json:"id" gorm:"primary_key"`
	StudentId    int       `json:"student_id" gorm:"default '0'; comment:'学员ID' "`
	Status       int       `json:"status" gorm:"default '0'; comment:'跟进状态' " `
	EmployeeId   int       `json:"employee_id" gorm:"default '0'; comment:'跟进人ID' "`
	PlanTime     string    `json:"plan_time" gorm:"type:char(10); not null ; default ''; comment:'计划时间' "`
	Result       string    `json:"result" gorm:"type:varchar(100); not null ; default ''; comment:'跟进结果' "`
	FollowerCon  string    `json:"follower_con" gorm:"type:varchar(100); not null ; default ''; comment:'跟进内容' "`
	FollowerImg  string    `json:"follower_img" gorm:"type:varchar(100); not null ; default ''; comment:'图片' "`
	FollowerType int       `json:"follower_type" gorm:"type:varchar(100); not null ; default ''; comment:'跟进方式' "`
	CompleteTime time.Time `json:"complete_time" time_format:"2006-01-02 15:04:05" gorm:"comment:'完成时间';not null ; default ''"`

	erm.Model
}
