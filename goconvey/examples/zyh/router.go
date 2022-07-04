package zyh

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)

	// 路径映射 ;
	/*userc := controller.NewUserController()
	router.GET("/user/get", userc.GetUser)
	router.POST("/user/post", userc.PostUser)
	*/
	return router
}

func HandleNotFound(c *gin.Context) {
	//global.NewResult(c).Error(404, "资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//global.NewResult(c).Error(500, "服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
