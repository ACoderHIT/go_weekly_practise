package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"snow.user/app/services/user"
	"strconv"
	"strings"
)

func Insert(c *gin.Context)  {
	mobile := c.PostForm("mobile")
	e ,err := user.Create(mobile)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}
	Success(c, "success" + string(e))
}

func MulInsert(c *gin.Context) {
	str := c.PostForm("mobile")
	if str == "" {
		Error(c, 500, "mobile参数必传，多个手机号以,分割")
		return
	}
	var mobiles []string
	for _, m := range strings.Split(str, ",") {
		mobiles = append(mobiles, m)
	}

	effected, err := user.MulCreate(mobiles)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	Success(c, fmt.Sprintf("成功写入：%d 条记录", effected))
}

func UpdateNameByMobile(c *gin.Context)  {
	mobile := c.PostForm("mobile")
	if mobile == "" {
		Error(c, 500, "mobile参数必传")
		return
	}

	name := c.PostForm("name")
	if name == "" {
		Error(c, 500, "name参数必传")
		return
	}

	effected, err := user.UpdateNameByMobile(mobile, name)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	Success(c, fmt.Sprintf("成功修改：%d 条记录", effected))
}

func UpdateMulNameById(c *gin.Context)  {
	idStr := c.PostForm("id")
	id, err := strconv.ParseInt(idStr, 0,64)
	if err != nil {
		Error(c, 500, "id必须是数字")
		return
	}
	if id <= 0 {
		Error(c, 500, "id参数必传")
		return
	}

	name := c.PostForm("name")
	if name == "" {
		Error(c, 500, "name参数必传")
		return
	}

	effected, err := user.UpdateMulNameById(id, name)
	if err != nil {
		Error(c, 500, err.Error())
		return
	}

	Success(c, fmt.Sprintf("成功修改：%d 条记录", effected))
}
