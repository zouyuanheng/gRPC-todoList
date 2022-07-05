package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	A int    `form:"a" binding:"-"`
	B string `from:"b" binding:"required"`
}

func main() {

	router := gin.Default()

	router.POST("register", Register)

	router.Run(":9999")
}

func Register(c *gin.Context) {
	var r RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"error_code": 21,
			"message":   "empty or wriong params",
			"reference": err.Error()})
		return
	}
	//验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, gin.H{"errir_code": 0,
		"message":   "success",
		"reference": "ok"})
}
