package entryTask

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetAccountWrongPath(t *testing.T) {
	r := gorequest.New()
	var UserName = "2"
	resp, body, errs := r.Post("http://127.0.0.1:4000/api/v1/user/register").Send(
		map[string]string{
			"UserName":        UserName,
			"NickName":        "1",
			"password":        "12345678",
			"PasswordConfirm": "12345678",
		},
	).End()
	if errs != nil {
		fmt.Println(errs)
	}

	Convey("测试登陆", t, func() {
		UserName = "1"
		resp.Body.Close()

		Convey("正确的code", func() {
			So(body, ShouldEqual, "{\"code\":404,\"msg\":\"userService--rpc error: code = Unknown desc = Error 1062: Duplicate entry '1' for key 'user.user_name'\"}")

		})

		UserName = "1"
		resp.Body.Close()
		Convey("错误的code", func() {
			So(body, ShouldEqual, "{\"cod1e\":404,\"msg\":\"userService--rpc error: code = Unknown desc = Error 1062: Duplicate entry '1' for key 'user.user_name'\"}")

		})
	})
}
