package handler

import (
	"todo/db"

	"github.com/gin-gonic/gin"
)

//增加借阅
func insertborrow(c *gin.Context)  {
	todo :=db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err  := todo.Insertborrow()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//删除借阅
func deleteborrow(c *gin.Context)  {
	todo :=db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Deleteborrow()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, todo)
}

//修改借阅
func updateborrow(c *gin.Context)  {
	todo :=db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Updateborrow()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询所有借阅
func selectborrows(c *gin.Context) {
	todo,err := db.Selectborrows()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询指定借阅
func selectborrow(c *gin.Context) {
	todo := db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectborrow()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//查询修改获取的用户
func selectupdateborrow(c *gin.Context)  {
	todo := db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectupdateborrow()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//返回修改用户的信息
func backupdateborrow(c *gin.Context)  {
	todo,err := db.Backupdateborrow()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询借阅人是否存在
func selectborrowuser(c *gin.Context)  {
	todo := db.Borrow_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectborrowuser()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}