package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 选择模式
	gin.SetMode(utils.AppMode)

	// 初始化路由
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("user", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口

		// 文章模块的路由接口
	}
	_ = r.Run(utils.HttpPort)
}