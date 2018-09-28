package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-ini/ini"
	"fmt"
	"time"
	"os"
)

var orm *gorm.DB

func init() {
	// load配置
	cfg, _ := ini.Load("conf/database.ini", "conf/app.ini")
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	// 主机
	host := cfg.Section(mode).Key("mysql.host").String()
	// 端口
	port := cfg.Section(mode).Key("mysql.port").String()
	// 用户名
	username := cfg.Section(mode).Key("mysql.username").String()
	// 密码
	password := cfg.Section(mode).Key("mysql.password").String()
	// 数据库名称
	dbname := cfg.Section(mode).Key("mysql.dbname").String()
	// 最大空闲连接数
	maxIdleConns, _ := cfg.Section(mode).Key("mysql.max_idle_conns").Int()
	// 最大打开的连接数
	maxOpenConns, _ := cfg.Section(mode).Key("mysql.max_open_conns").Int()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	var err error
	orm, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Fail to open mysql: %v", err)
		os.Exit(1)
	}

	orm.DB().SetMaxIdleConns(maxIdleConns)
	orm.DB().SetMaxOpenConns(maxOpenConns)
	orm.DB().SetConnMaxLifetime(time.Hour)
}

func GetORM() *gorm.DB {
	return orm
}