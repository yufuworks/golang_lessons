package s8

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port       int
	DbName     string
	SQLDriver  string
	SQLDriver2 string
}

// iniファイルを読み込んで、指定した構造体に当てる
var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:       cfg.Section("web").Key("port").MustInt(),
		DbName:     cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		SQLDriver2: cfg.Section("db").Key("driver2").String(),
	}
}

func Main06() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
