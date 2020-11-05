package router

import (
	"gin/controller"
	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()
	//添加用户
	e.POST("/addUser", controller.AddUser)
	//获取所有用户
	e.GET("/user/getUsers", controller.GetUsers)
	//更新用户
	e.POST("/user/updateUser", controller.UpdateUser)
	//删除单个用户
	e.GET("/user/del/:id", controller.DeleteUserByID)
	//获取单个用户
	e.GET("/user/get/:id", controller.GetUserByID)

	e.Start(":8080")
}
