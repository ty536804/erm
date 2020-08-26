package Services

import (
	"erm/Model/Admin"
	"erm/Pkg/Common"
	"erm/Pkg/E"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summer 登录校验
func Login(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return E.ERROR, err.Error()
	}

	uName := Common.GetStringVal(c.PostForm("uname"))
	pwd := Common.GetStringVal(c.PostForm("pword"))

	valid := validation.Validation{}
	valid.Required(uName, "uname").Message("用户名不能为空")
	valid.Required(pwd, "pword").Message("密码不能为空")

	if !valid.HasErrors() {
		AdminInfo := Admin.GetAdmin(uName)
		if AdminInfo.Id < 1 {
			return E.ERROR, "用户不存在"
		}
		if AdminInfo.Pwd != Common.Md5Pwd(pwd) {
			return E.ERROR, "用户活密码不正确"
		}
		_, RES := Common.GernToken(c, AdminInfo.Id)
		return E.SUCCESS, RES
	}
	return Common.ValidErr(valid)
}

// @Summer 添加管理员
func AddAdmin(c *gin.Context) (code int, msg string) {
	var data = make(map[string]interface{})
	if err := c.Bind(&c.Request.Body); err != nil {
		return E.ERROR, err.Error()
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

		isOk := true
		if id < 1 {
			isOk = Admin.AddAdminUser(data)
		} else {
			isOk = Admin.EditAdminUser(id, data)
		}
		if isOk {
			return E.SUCCESS, "操作成功"
		}
		return E.ERROR, "操作失败"
	}
	return Common.ValidErr(valid)
}

// @Summer 获取权限列表
// @Param total int 总计
// @Param list  []  权限列表
func GetAdmins(c *gin.Context) (res map[string]interface{}) {
	var data = make(map[string]interface{})
	page := com.StrTo(c.PostForm("page")).MustInt()
	data["list"] = Admin.GetAdminUserList(page, data)
	data["total"] = Admin.GetTotalAdmin()
	return data
}

// @Summer 通过制定ID获取管理员信息
func GetAdmin(c *gin.Context) (data map[string]interface{}) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	res := make(map[string]interface{})
	res["user"] = Admin.GetAdminUser(id)
	return res
}
