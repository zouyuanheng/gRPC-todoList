package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	UserName        string `form:"UserName" binding:"required"`
	NickName        string `form:"NickName" binding:"required"`
	Password        string `form:"Password" binding:"required"`
	PasswordConfirm string `form:"PasswordConfirm" binding:"required"`
}

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

/*func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.ShouldBindBodyWith(&userReq, binding.JSON))
	var r RegisterRequest
	err := ginCtx.ShouldBindBodyWith(&r, binding.JSON) //.(service.UserServiceClient)
	if err != nil {
		fmt.Println("register failed")
		ginCtx.JSON(http.StatusOK, gin.H{"error_code": 21,
			"message":   "empty or wriong params",
			"reference": err.Error()})
		return
	}

	fmt.Println("register success")
	ginCtx.JSON(http.StatusOK, gin.H{"error_code": 0,
		"message":   "success",
		"reference": "ok"})

}*/

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Data:   res.TokenData{User: userResp.UserDetail, Token: token},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
