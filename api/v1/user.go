package v1

import (
	"ginblog/model"
	"ginblog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 状态码
var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// CheckUser() 会返回一个判断用户是否存在的状态码
	code = model.CheckUser(data.UserName)
	if code == err_msg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == err_msg.ErrorUsernameUsed {
		code = err_msg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": err_msg.GetErrMsg(code),
	})

}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code = err_msg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": err_msg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == err_msg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == err_msg.ErrorUsernameUsed {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": err_msg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": err_msg.GetErrMsg(code),
	})
}
