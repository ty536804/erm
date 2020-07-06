package Services

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"github.com/unknwon/com"
)

func ReadExcel(c *gin.Context) {

	valid := validation.Validation{}

	if err := c.Bind(&c.Request.Body); err != nil {
		fmt.Println(err)
	}

	file := com.StrTo(c.PostForm("file")).String()
	valid.Required(file, "file").Message("必须上传文件")
	xlsx.OpenFile(file)
}
