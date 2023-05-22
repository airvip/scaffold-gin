package config

// OSS 配置
type OssConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	BasePath        string
}

// sms 阿里
type SmsAliConfig struct {
	AccessKeyId     string
	AccessKeySecret string
}
