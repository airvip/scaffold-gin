package config

// mysql 配置
type DbConfig struct {
	User    string
	Pass    string
	Host    string
	Port    string
	DbName  string
	Charset string
}
