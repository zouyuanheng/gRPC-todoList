package examples

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"testing"
)

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request 1", t, func() {
		req, err := http.Get("127.0.0.1:4000/api/v1/task")
		if err != nil {
			log.Println(err)
			return
		}

		Convey("When the request is handled by the Router", func() {

			Convey("Then the response should be a 404", func() {
				So(req.Body, ShouldEqual, 404)
			})
		})
	})
}
