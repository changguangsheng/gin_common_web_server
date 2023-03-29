package initialize

import (
	"ewa_admin_server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgsql()
	default:
		return GormMysql()
	}
}
