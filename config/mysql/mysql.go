package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

// Init 初始化mysql数据库
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
		true,
		//"Asia/Shanghai"),
		"Local")
	zap.L().Info(dsn)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("database connection failed, err:%v", err))
	}

	setupDB(Db)

	return
}

// Close 关闭数据库连接
func Close() {
	sqlDB, err := Db.DB()
	if err != nil {
		panic(fmt.Errorf("get sqlDB failed, err:%v", err))
	}
	sqlDB.Close()
}

// mysql 相关配置
func setupDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("get sqlDB failed, err:%v", err))
	}
	sqlDB.SetMaxOpenConns(viper.GetInt("db.max_open_conns")) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxIdleConns(viper.GetInt("db.max_idle_conns")) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}
