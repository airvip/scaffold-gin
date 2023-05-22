package global

import (
	"scaffold-gin/common/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OSS = InitOss()

func InitOss() *oss.Bucket{

	bucketName := config.Conf.OSS.BucketName
	endpoint := config.Conf.OSS.Endpoint
	accessKeyId := config.Conf.OSS.AccessKeyId
	accessKeySecret := config.Conf.OSS.AccessKeySecret

    client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
    if err != nil {
		ZAPLOGGER.Error("oss Init Error, "+err.Error())
		return nil
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		ZAPLOGGER.Error(bucketName + " is Error, "+err.Error())
		return nil
	}
	return bucket
}
