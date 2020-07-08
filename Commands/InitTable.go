package Commands

import (
	erm "erm/Database"
	"erm/Model/Admin"
	"erm/Model/Attends"
	"erm/Model/Courses"
	"erm/Model/Rules"
	"erm/Model/Sales"
	"erm/Model/School"
	"erm/Model/Users"
)

func init() {
	GenerateDatabase()
}

func GenerateDatabase() {
	erm.Db.AutoMigrate(&Admin.SysAdminUser{})
	erm.Db.AutoMigrate(&Admin.SysAdminPower{})
	erm.Db.AutoMigrate(&Admin.SysAdminDepartment{})

	erm.Db.AutoMigrate(&Attends.AttendLog{})

	erm.Db.AutoMigrate(&Courses.Schedule{})
	erm.Db.AutoMigrate(&Courses.CourseMethod{})
	erm.Db.AutoMigrate(&Courses.Course{})

	erm.Db.AutoMigrate(&Rules.Rule{})

	erm.Db.AutoMigrate(&Sales.Student{})
	erm.Db.AutoMigrate(&Sales.Listen{})
	erm.Db.AutoMigrate(&Sales.Follower{})

	erm.Db.AutoMigrate(&School.ClassName{})
	erm.Db.AutoMigrate(&School.Employee{})
	erm.Db.AutoMigrate(&School.School{})

	erm.Db.AutoMigrate(&Users.User{})
}
func DropDatabase() {
	DropAdmin()
	DropAttends()
	DropCourses()
	DropRule()
	DropSales()
	DropSchool()
	DropUsers()
}

// @Summer 管理员表
func DropAdmin() {
	if !erm.Db.HasTable(&Admin.SysAdminUser{}) {
		erm.Db.DropTable(&Admin.SysAdminUser{})
	}

	if !erm.Db.HasTable(&Admin.SysAdminPower{}) {
		erm.Db.DropTable(&Admin.SysAdminPower{})
	}

	if !erm.Db.HasTable(&Admin.SysAdminDepartment{}) {
		erm.Db.DropTable(&Admin.SysAdminDepartment{})
	}
}

// 点名记录
func DropAttends() {
	if !erm.Db.HasTable(&Attends.AttendLog{}) {
		erm.Db.DropTable(&Attends.AttendLog{})
	}
}

func DropCourses() {
	if !erm.Db.HasTable(&Courses.CourseMethod{}) {
		erm.Db.DropTable(&Courses.CourseMethod{})
	}

	if !erm.Db.HasTable(&Courses.Course{}) {
		erm.Db.DropTable(&Courses.Course{})
	}

	if !erm.Db.HasTable(&Courses.Schedule{}) {
		erm.Db.DropTable(&Courses.Schedule{})
	}
}

func DropRule() {
	if !erm.Db.HasTable(&Rules.Rule{}) {
		erm.Db.DropTable(&Rules.Rule{})
	}
}

func DropSales() {
	if !erm.Db.HasTable(&Sales.Follower{}) {
		erm.Db.DropTable(&Sales.Follower{})
	}

	if !erm.Db.HasTable(&Sales.Listen{}) {
		erm.Db.DropTable(&Sales.Listen{})
	}

	if !erm.Db.HasTable(&Sales.Student{}) {
		erm.Db.DropTable(&Sales.Student{})
	}
}

func DropSchool() {
	if !erm.Db.HasTable(&School.Employee{}) {
		erm.Db.DropTable(&School.Employee{})
	}

	if !erm.Db.HasTable(&School.School{}) {
		erm.Db.DropTable(&School.School{})
	}

	if !erm.Db.HasTable(&School.ClassName{}) {
		erm.Db.DropTable(&School.ClassName{})
	}
}

func DropUsers() {
	if !erm.Db.HasTable(&Users.User{}) {
		erm.Db.DropTable(&Users.User{})
	}
}
