package handler

import (
	"todo/db"

	"github.com/gin-gonic/gin"
)

//增加用户
func insertuser(c *gin.Context)  {
	todo :=db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err  := todo.Insertuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//删除用户
func deleteuser(c *gin.Context)  {
	todo :=db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Deleteuser()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, todo)
}

//修改用户
func updateuser(c *gin.Context)  {
	todo :=db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Updateuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询所有用户
func selectusers(c *gin.Context) {
	todo,err := db.Selectusers()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询指定用户
func selectuser(c *gin.Context) {
	todo := db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}
//登录查询用户是否存在
func selectloginuser(c *gin.Context)  {
	todo := db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectloginuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}
//查询修改获取的用户
func selectupdateuser(c *gin.Context)  {
	todo := db.User_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectupdateuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//返回修改用户的信息
func backupdateuser(c *gin.Context)  {
	todo,err := db.Backupdateuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}