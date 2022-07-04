package entryTask

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:4000/api/v1/user/register").Send(
			map[string]string{
				"UserName":        "1",
				"NickName":        "1",
				"password":        "12345678",
				"PasswordConfirm": "12345678",
			},
		).End()
		if errs != nil {
			fmt.Println(errs)
		}
		Convey("When the request is handled by the Router", func() {
			resp.Body.Close()

			Convey("Then the response should be a 404", func() {
				So(body, ShouldEqual, "{\"cod1e\":404,\"msg\":\"userService--rpc error: code = Unknown desc = Error 1062: Duplicate entry '1' for key 'user.user_name'\"}")
			})
		})
	})
}
