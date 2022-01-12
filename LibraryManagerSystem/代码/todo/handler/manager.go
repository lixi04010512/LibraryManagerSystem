package handler

import (
	"todo/db"

	"github.com/gin-gonic/gin"
)

//增加管理员
func insertmanager(c *gin.Context)  {
	todo :=db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err  := todo.Insertmanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//删除管理员
func deletemanager(c *gin.Context)  {
	todo :=db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Deletemanager()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, todo)
}

//修改管理员
func updatemanager(c *gin.Context)  {
	todo :=db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Updatemanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询所有管理员
func selectmanagers(c *gin.Context) {
	todo,err := db.Selectmanagers()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询指定管理员
func selectmanager(c *gin.Context) {
	todo := db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectmanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}
//登录查询管理员是否存在
func selectloginmanager(c *gin.Context)  {
	todo := db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectloginmanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}
//查询修改获取的管理员
func selectupdatemanager(c *gin.Context)  {
	todo := db.Manager_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectupdatemanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//返回修改管理员的信息
func backupdatemanager(c *gin.Context)  {
	todo,err := db.Backupdatemanager()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}