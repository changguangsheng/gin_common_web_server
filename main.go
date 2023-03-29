package main

import (
	"ewa_admin_server/core"
	"ewa_admin_server/global"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	/*
		DebugMode比ReleaseMode多了一些额外的错误信息，生产环境不需要这些信息。
		而TestMode是gin用于自己的单元测试，用来快速开关DebugMode。对其它开发者没什么意义。
		可以通过gin.SetMode(AppMode)来设置mode。

	*/
	gin.SetMode(AppMode)

	// TODO：1.配置初始化
	global.EWA_VIPER = core.InitializeViper()

	// TODO：2.日志
	global.EWA_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.EWA_LOG)

	global.EWA_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	//  TODO：3.数据库连接

	// TODO：4.其他初始化

	// TODO：5.启动服务
	core.RunServer()
}
