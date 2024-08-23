package app

import (
	"house/app/model"
	"house/app/router"
)

// Start 这是一个启动器方法
func Start() {
	// 初始化数据库和其他服务
	model.NewMysql()
	model.NewRdb()

	// 启动路由器
	router.New()

	// 在程序结束时关闭数据库连接等资源
	defer func() {
		model.Close()
	}()
}
