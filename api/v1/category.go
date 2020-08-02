package v1

import (
	"ginblog/model"
	"ginblog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// CheckUser() 会返回一个判断用户是否存在的状态码
	code = model.CheckCategory(data.Name)
	if code == err_msg.SUCCESS {
		model.CreateCate(&data)
	}
	if code == err_msg.ErrorCategoryNameUsed {
		code = err_msg.ErrorCategoryNameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": err_msg.GetErrMsg(code),
	})

}

// 查询单个分类下的文章

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCate(pageSize, pageNum)
	code = err_msg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": err_msg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == err_msg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == err_msg.ErrorCategoryNameUsed {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": err_msg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": err_msg.GetErrMsg(code),
	})
}
