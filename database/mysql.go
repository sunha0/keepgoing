package database

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/keepgoing/global"
	"github.com/keepgoing/models/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormMysql() *gorm.DB {
	m := global.Conf.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Conf.Mysql.Prefix,
			SingularTable: global.Conf.Mysql.Singular,
		},
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 迁移mysql数据库表
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		&system.SysUser{},
		&system.SysRole{},
		&system.SysUserRole{},

	)
	if err != nil {
		// 	global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	// global.GVA_LOG.Info("register table success")
}
