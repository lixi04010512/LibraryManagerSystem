package handler

import (
	"fmt"
	"todo/db"

	"github.com/gin-gonic/gin"
)

//增加图书
func insertbook(c *gin.Context)  {
	todo :=db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err  := todo.Insertbook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//删除图书
func deletebook(c *gin.Context)  {
	todo :=db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Deletebook()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, todo)
}

//修改图书
func updatebook(c *gin.Context)  {
	todo :=db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Updatebook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//补给图书
func addbook(c *gin.Context)  {
	todo :=db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	err :=todo.Addbook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询所有图书
func selectbooks(c *gin.Context) {
	todo,err := db.Selectbooks()
	if err != nil {
		respError(c, err)
		return
	}
    fmt.Println(todo)
	respOK(c, todo)
}

//查询指定图书
func selectbook(c *gin.Context) {
	todo := db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectbook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//查询修改获取的图书
func selectupdatebook(c *gin.Context)  {
	todo := db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectupdatebook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//返回修改图书的信息
func backupdatebook(c *gin.Context)  {
	todo,err := db.Backupdatebook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}

//查询修改获取的图书
func selectaddbook(c *gin.Context)  {
	todo := db.Book_lixi{}
	if err := c.ShouldBind(&todo); err != nil {
		respError(c, err)
		return
	}
	todo1,err  := todo.Selectaddbook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo1)
}

//返回修改图书的信息
func backaddbook(c *gin.Context) {
	todo, err := db.Backaddbook()
	if err != nil {
		respError(c, err)
		return
	}

	respOK(c, todo)
}