package Sales

import (
	erm "erm/Database"
	"time"
)

// @Summer 试听记录
type Listen struct {
	Id          int       `json:"id" gorm:"primary_key"`
	StudentId   int       `json:"student_id" gorm:"default '0'; comment:'学员ID' "`
	CourseId    int       `json:"course_id" gorm:"comment:'课程ID'"`
	SClassId    int       `json:"course_id" gorm:"comment:'班级ID'"`
	ReserveTime time.Time `json:"reserve_time" time_format:"2006-01-02 15:04:05" gorm:"not null default: '';comment:'预约时间'" `
	LStatus     int       `json:"l_status" gorm:"default:'1';comment:'状态 0已取消 1已预约'" `

	erm.Model
}
