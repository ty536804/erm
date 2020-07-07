package Commands

import (
	"erm/Model/Admin"
	"erm/Pkg/Common"
)

func init() {
	//AddUser()
}

//初始化管理员信息
func AddUser() {
	var data = make(map[string]interface{})
	if !Admin.ExistsByLoginName("admin") {
		data["nick_name"] = "admin"
		data["pwd"] = Common.Md5Pwd("admin")
		data["login_name"] = "admin"
		data["email"] = "admin@126.com"
		data["status"] = 1
		data["tel"] = "13643513445"
		data["city_id"] = 10000
		data["avatar"] = Common.GetDir() + "/Resources/Public/upload/avator.gif"
		data["department_id"] = ""
		data["power_id"] = ""
		Admin.AddAdminUser(data)
	}
}
