package examples

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func rw111() {
	resp, err := http.Get("http://127.0.0.1:4000/api/v1/task")

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header
	// headers 打印报文头部信息
	for k, v := range headers {
		fmt.Printf("%v, %v\n", k, v) // %v 打印interfac{}的值
	}

	// 打印响应信息内容
	fmt.Printf("响应状态：%s,响应码： %d\n", resp.Status, resp.StatusCode)
	fmt.Printf("协议：%s\n", resp.Proto)
	fmt.Printf("响应内容长度： %d\n", resp.ContentLength)
	fmt.Printf("编码格式：%v\n", resp.TransferEncoding) // 未指定时为空
	fmt.Printf("是否压缩：%t\n", resp.Uncompressed)
	fmt.Println(reflect.TypeOf(resp.Body)) // *http.gzipReader
	fmt.Println(resp.Close)

	buf := bytes.NewBuffer(make([]byte, 0, 512))
	length, _ := buf.ReadFrom(resp.Body)
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	fmt.Println(string(buf.Bytes()))
}
