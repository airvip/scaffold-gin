package config

// 应用级别的 zaplog - logger 配置
type LoggerConfig struct {
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Level      string
}
