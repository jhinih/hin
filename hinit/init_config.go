package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/spf13/viper"
	"time"
)

func InitConfig() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", hglobal.Path+hglobal.DEFAULT_CONFIG_FILE_PATH, "配置文件绝对路径或相对路径")
	flag.Parse()
	configPath = "./config/config.yaml"
	fmt.Printf("配置文件路径为 %s\n", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件发生变化\n")
		if err := viper.Unmarshal(&hconfig.Conf); err != nil {
			fmt.Printf("无法反序列化配置文件 %v\n", err)
		}
		fmt.Printf("%+v\n", hconfig.Conf)
		hglobal.Config = hconfig.Conf
		Eve()
		Init()
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("无法读取配置文件 err: %v\n", err)

	}
	// 解析到变量中
	if err := viper.Unmarshal(&hconfig.Conf); err != nil {
		fmt.Printf("无法解析配置文件 err: %v\n", err)
	}
	hglobal.Config = hconfig.Conf
}
