package controller

import (
	"fmt"
	"gin/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)
/*
 *description：添加用户
 */
func AddUser(ctx echo.Context) error {
	name := ctx.FormValue("name")
	if name == "" {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "名字不可以为空",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	password := ctx.FormValue("password")
	if password == "" {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "密码不可以为空",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	phone := ctx.FormValue("phone")
	if phone == "" {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "电话号码不可以为空",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	user2 := &model.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	fmt.Println(user2)
	result := model.AddUser(user2)

	if result != nil {
		rs := map[string]interface{}{
			"data":    "",
			"code":    "5",
			"message": "添加失败",
		}
		return ctx.JSON(http.StatusOK,rs)
	}else{
		rs := map[string]interface{}{
			"data":   user2 ,
			"code":    "2",
			"message": "添加成功",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
}
/*
 *description：获取所有用户
 */
func GetUsers(ctx echo.Context) error {
	users, err := model.GetUsers()
	if err != nil {
		rs := map[string]interface{}{
			"data":    "",
			"code":    "5",
			"message": "查询失败",
		}
		return ctx.JSON(http.StatusOK,rs)

	} else {
		rs := map[string]interface{}{
			"data":    users,
			"code":    "2",
			"message": "查询成功",
		}
		return ctx.JSON(http.StatusOK,rs)
	}
}
/*
 *description：更新用户信息
 */
func UpdateUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "请传入正确的ID",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	name := ctx.FormValue("name")
	password := ctx.FormValue("password")
	phone := ctx.FormValue("phone")
	user := &model.User{
		ID:id,
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	err2 := model.UpdateUser(user)
	if err2 != nil {
		rs := map[string]interface{}{
			"data":    user,
			"code":    "5",
			"message": "更新失败",
		}
		return ctx.JSON(http.StatusOK,rs)
	}else{
		rs := map[string]interface{}{
			"data":    user,
			"code":    "2",
			"message": "更新成功",
		}
		return ctx.JSON(http.StatusOK,rs)
	}
}
/*
 *description：删除单个用户
 */
func DeleteUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "请传入正确的ID",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	err2 := model.DeleteUserByID(id)
	if err2 != nil {
		rs := map[string]interface{}{
			"data":    "",
			"code":    "5",
			"message": "删除失败",
		}
		return ctx.JSON(http.StatusOK,rs)
	}else{
		rs := map[string]interface{}{
			"data":    "",
			"code":    "2",
			"message": "删除成功",
		}
		return ctx.JSON(http.StatusOK,rs)
	}
}
/*
 *description：获取单个用户
 */
func GetUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		rs := map[string]interface{}{
			"data":   "" ,
			"code":    "3",
			"message": "请传入正确的ID",
		}
		return ctx.JSON(http.StatusOK, rs)
	}
	users, err1 := model.GetUserByID(id)
	if err1 != nil {
		rs := map[string]interface{}{
			"data":    "",
			"code":    "5",
			"message": "单条查找失败",
		}
		return ctx.JSON(http.StatusOK,rs)
	}else{
		rs := map[string]interface{}{
			"data":    users,
			"code":    "2",
			"message": "单条查找成功",
		}
		return ctx.JSON(http.StatusOK,rs)
	}
}