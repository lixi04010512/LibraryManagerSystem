package handler

import (
	"github.com/gin-gonic/gin"
)

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()

	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器
		r.Static("/", webDir)
	}

	// api接口服务，定义了路由组 /todo
	todo := r.Group("/todo")
	{
		// 定义增改查的接口，并注册到web服务器

		//管理员
		todo.POST("/selectmanagers", selectmanagers)
		todo.POST("/selectmanager", selectmanager)
		todo.POST("/selectloginmanager",selectloginmanager)
		todo.POST("/selectupdatemanager",selectupdatemanager)
		todo.POST("/backupdatemanager",backupdatemanager)
		todo.POST("/insertmanager", insertmanager)
		todo.POST("/deletemanager", deletemanager)
		todo.POST("/updatemanager",updatemanager)

		//用户
		todo.POST("/selectusers", selectusers)
		todo.POST("/selectuser", selectuser)
		todo.POST("/selectloginuser",selectloginuser)
		todo.POST("/selectupdateuser",selectupdateuser)
		todo.POST("/selectborrowuser",selectborrowuser)
		todo.POST("/backupdateuser",backupdateuser)
		todo.POST("/insertuser", insertuser)
		todo.POST("/deleteuser", deleteuser)
		todo.POST("/updateuser",updateuser)

		//图书
		todo.POST("/selectbooks",selectbooks)
		todo.POST("/selectbook", selectbook)
		todo.POST("/selectupdatebook",selectupdatebook)
		todo.POST("/backupdatebook",backupdatebook)
		todo.POST("/selectaddbook",selectaddbook)
		todo.POST("/backaddbook",backaddbook)
		todo.POST("/insertbook", insertbook)
		todo.POST("/addbook", addbook)
		todo.POST("/deletebook", deletebook)
		todo.POST("/updatebook",updatebook)

		//借阅
		todo.POST("/selectborrows", selectborrows)
		todo.POST("/selectborrow", selectborrow)
		todo.POST("/selectupdateborrow",selectupdateborrow)
		todo.POST("/backupdateborrow",backupdateborrow)
		todo.POST("/insertborrow", insertborrow)
		todo.POST("/deleteborrow", deleteborrow)
		todo.POST("/updateborrow",updateborrow)

	}

	// 启动web服务
	err = r.Run(addr)
	return err
}

// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}

