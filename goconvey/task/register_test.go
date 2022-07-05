package task

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//
func TestHttp(t *testing.T) {
	Convey("参数“a”和参数“b”传的值都正确", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{
			"a": 2,
			"b": "2"}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "success")
	})
	Convey("参数“a“传字符串", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{
			"a": "2",
			"b": "2"}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "json: cannot unmarshal string into Go struct field RegisterRequest.A of type int")
	})
	Convey("参数“b“传整型", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{
			"a": 2,
			"b": 2}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "json: cannot unmarshal number into Go struct field RegisterRequest.B of type string")
	})
	Convey("参数“b”没有填值", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{
			"a": 2,
		}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "Key: 'RegisterRequest.B' Error:Field validation for 'B' failed on the 'required' tag")
	})
	Convey("参数“a”没有填值", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{
			"b": "2",
		}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "success")
	})
	Convey("参数“a”和参数“b”传的值都为空", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{"c": 2}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "Key: 'RegisterRequest.B' Error:Field validation for 'B' failed on the 'required' tag")
	})

	Convey("执行失败的测试案例", t, func() {
		r := gorequest.New()
		resp, body, errs := r.Post("http://127.0.0.1:9999/register").Send(map[string]interface{}{"c": 2}).End()
		if errs != nil {
			fmt.Println(errs)
		}
		resp.Body.Close()

		So(body, ShouldContainSubstring, "success")
	})

}
