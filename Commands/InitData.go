package Commands

import (
	erm "erm/Database"
	"erm/Model/Admin"
	"erm/Pkg/Common"
	"fmt"
)

func init() {
	//AddUser()
	AddPower()
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

func AddPower() {
	erm.Db.First(&Admin.SysAdminPower{})
	//教务中心
	sql := "INSERT INTO `sys_admin_power` (`power_name`,`level`,`pid`,`status`,`desc`) VALUES ('教务中心','0','0','1','教务中心')," +
		"('学员管理','1','1','1','学员管理'),('班级管理','1','1','1','班级管理'),('课表管理','1','1','1','课表管理'),('上课记录','1','1','1','上课记录')," +
		"('老师管理','1','1','1','老师管理')," +
		"('课程管理','1','1','1','课程管理'),('物品/费用','1','1','1','物品/费用'),('在线商城','1','1','1','在线商城'),('人脸考勤','1','1','1','人脸考勤')," +
		"('刷卡考勤','1','1','1','刷卡考勤')"

	result := erm.Db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("数据填充失败:%s", result.Error)
	}

	//家校互动
	sql = "INSERT INTO `sys_admin_power` (`power_name`,`level`,`pid`,`status`,`desc`) VALUES ('电子相册','0','0','1','电子相册')," +
		"('成长答案','1','12','1','成长答案'),('课后点评','1','12','1','课后点评'),('课后作业','1','12','1','课后作业'),('成绩单','1','12','1','成绩单')," +
		"('通知管理','1','12','1','通知管理')"

	result = erm.Db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("数据填充失败:%s", result.Error)
	}

	//销售中心
	sql = "INSERT INTO `sys_admin_power` (`power_name`,`level`,`pid`,`status`,`desc`) VALUES ('销售中心','0','0','1','销售中心')," +
		"('线索管理','1','18','1','线索管理'),('试听记录','1','18','1','试听记录'),('跟进管理','1','18','1','跟进管理'),('续费缴费','1','18','1','续费缴费')"

	result = erm.Db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("数据填充失败:%s", result.Error)
	}

	//财务中心
	sql = "INSERT INTO `sys_admin_power` (`power_name`,`level`,`pid`,`status`,`desc`) VALUES ('财务中心','0','0','1','财务中心')," +
		"('订单管理','1','23','1','订单管理'),('收支明细','1','23','1','收支明细'),('工资结算','1','23','1','工资结算'),('课消记录','1','23','1','课消记录')," +
		"('确认收入报表','1','23','1','确认收入报表'),('账户充值','1','23','1','账户充值')"

	result = erm.Db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("数据填充失败:%s", result.Error)
	}

	//数据中心
	sql = "INSERT INTO `sys_admin_power` (`power_name`,`level`,`pid`,`status`,`desc`) VALUES ('数据中心','0','0','1','数据中心')," +
		"('销售数据','1','30','1','销售数据'),('教务数据','1','30','1','教务数据'),('家校数据','1','30','1','家校数据'),('财务数据','1','30','1','财务数据')"
	result = erm.Db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("数据填充失败:%s", result.Error)
	}

}
