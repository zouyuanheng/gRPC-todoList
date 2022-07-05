#！/bin/bash


#1、启动etcd
etcd
#2、启动api-gateway
cd /Users/yuanheng.zou/Desktop/entryTask/api-gateway
go run cmd/main.go
#3、启动task服务
cd /Users/yuanheng.zou/Desktop/entryTask/task
go run cmd/main.go
#4、启动user服务
cd /Users/yuanheng.zou/Desktop/entryTask/user
go run cmd/main.go
#5、启动apiForTest服务
cd /Users/yuanheng.zou/Desktop/entryTask/goconvey
go run apiForTest.go

