package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Runmode      string `mapstructure:"runmode"`
	Addr         string `mapstructure:"addr"`
	*MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Port         string `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	Db           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init() error {
	viper.SetConfigFile("config/conf/config.yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper ReadInConfig failed, err:%v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unmarshal to conf failed, err:%v", err))
	}
	return err
}
