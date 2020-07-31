package v1

import (
	"ginblog/model"
	"ginblog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
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
	code = model.CheckUser(data.UserName);
	if code == err_msg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == err_msg.ErrorUsernameUsed {
		code = err_msg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": err_msg.GetErrMsg(code),
	})

}
// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {

}
// 编辑用户
func EditUser(c *gin.Context) {

}
// 删除用户
func DeleteUser(c *gin.Context) {

}