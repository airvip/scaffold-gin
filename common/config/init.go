package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	// configPath = "./etc/app.yml"
)

// 全局配置，需要哪个添加哪个
type Config struct {
	Server ServiceConfig
	Logger LoggerConfig // 应用级别的 zaplog
	
	DB     DbConfig
	REDIS  RedisConfig
    // ali series
	OSS    OssConfig
	SmsAli SmsAliConfig

	// tx series
	SmsTx  SmsTxConfig
}

var Conf = &Config{}

// func init() {
func InitConfig(configPath string) {

	viper.SetConfigFile(configPath)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})

}
