package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/keepgoing/config"
	"github.com/keepgoing/global"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值

func Viper(path ...string) *viper.Viper {
	var conf string

	if len(path) == 0 {
		flag.StringVar(&conf, "c", "", "choose config file.")
		flag.Parse()

		if conf == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" { // 判断 config.ConfigEnv 常量存储的环境变量是否为空

				switch gin.Mode() {
				case gin.DebugMode:
					conf = config.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigDefaultFile)
				case gin.ReleaseMode:
					conf = config.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigReleaseFile)
				case gin.TestMode:
					conf = config.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config.ConfigTestFile)
				}
			} else { // config.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				conf = configEnv

				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", config.ConfigEnv, conf)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", conf)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		conf = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", conf)
	}

	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Conf); err != nil {
			fmt.Println(err)
		}
	})

	// 序列化到config配置结构体，可以通过结构体调用参数
	if err = v.Unmarshal(&global.Conf); err != nil {
		panic(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// global.CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
