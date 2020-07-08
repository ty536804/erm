package Courses

import (
	erm "erm/Database"
)

// @Summer class schedule list
type Schedule struct {
	Id           int    `json:"id" gorm:"primary_key:true;unique"`
	ScheduleName string `json:"course_name" gorm:"type:varchar(100); not null default: '';comment:'课时'" `
	CourseId     int    `json:"course_id" gorm:"type:varchar(100); not null default: '';comment:'课程ID'"`

	erm.Model
}
