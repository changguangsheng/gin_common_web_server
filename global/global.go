package global

import (
	"ewa_admin_server/config"
	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
	//将 *gorm.DB 挂载到全局变量中的原因是为了方便在整个应用程序中共享一个数据库连接
	EWA_DB *gorm.DB
)
