package Services

import (
	"erm/Pkg/Common"
	"erm/Pkg/E"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddPower(c *gin.Context) (code int, msg string) {
	if err := c.Bind(&c.Request.Body); err != nil {
		return E.ERROR, err.Error()
	}

	powerName := Common.GetStringVal(c.PostForm("power_name"))
	level := Common.GetIntVal(c.PostForm("level"))
	pid := Common.GetIntVal(c.PostForm("pid"))
	status := Common.GetIntVal(c.PostForm("status"))
	desc := Common.GetStringVal(c.PostForm("desc"))
	fmt.Println(powerName, level, pid, status, desc)
	return E.ERROR, "111"
}
