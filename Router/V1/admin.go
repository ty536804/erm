package V1

import (
	"erm/Model/Admin"
	"erm/Pkg/Common"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 添加/编辑用户
func AddAdmin(c *gin.Context) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		Common.Error(c, err.Error(), data)
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	loginName := com.StrTo(c.PostForm("login_name")).String()
	pwd := com.StrTo(c.PostForm("pwd")).String()
	status := com.StrTo(c.PostForm("status")).MustInt()
	cityId := com.StrTo(c.PostForm("city_id")).MustInt()
	avatar := com.StrTo(c.PostForm("avatar")).String()
	tel := com.StrTo(c.PostForm("tel")).String()
	email := com.StrTo(c.PostForm("email")).String()
	nickName := com.StrTo(c.PostForm("nick_name")).String()
	departmentId := com.StrTo(c.PostForm("department_id")).MustInt()
	positionId := com.StrTo(c.PostForm("position_id")).String()
	schoolName := com.StrTo(c.PostForm("school_name")).String()

	valid := validation.Validation{}

	valid.Required(loginName, "login_name").Message("登录账号不能为空")
	valid.Required(pwd, "pwd").Message("密码不能为空")
	valid.Required(cityId, "city_id").Message("城市不能为空")
	valid.Required(tel, "tel").Message("手机号码不能为空")
	if err, msg := Common.CheckMobile(tel); err != nil {
		Common.Error(c, msg, data)
	}
	if !valid.HasErrors() {
		data["login_name"] = loginName
		data["pwd"] = pwd
		data["status"] = status
		data["city_id"] = cityId
		data["avatar"] = avatar
		data["tel"] = tel
		data["email"] = email
		data["nick_name"] = nickName
		data["department_id"] = departmentId
		data["position_id"] = positionId
		data["school_name"] = schoolName

		isOk := true
		if id < 1 {
			isOk = Admin.AddAdminUser(data)
		} else {
			isOk = Admin.EditAdminUser(id, data)
		}
		if isOk {
			Common.Success(c, "操作成功", "")
		}
		Common.Error(c, "操作失败", "")
	}
	Common.ViewErr(c, valid)
}

// @Summer 获取权限列表
// @Param total int 总计
// @Param list  []  权限列表
func GetAdmin(c *gin.Context) {
	var data = make(map[string]interface{})
	page := com.StrTo(c.PostForm("page")).MustInt()
	data["list"] = Admin.GetAdminUserList(page, data)
	data["total"] = Admin.GetTotalAdmin()
	Common.Success(c, "操作成功", data)
}
