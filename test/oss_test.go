package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 上传Byte数组
func TestOss(t *testing.T) {
	client, err := oss.New("yourEndpoint", "yourAccessKeyId", "yourAccessKeySecret")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 填写存储空间名称，例如examplebucket。
    bucket, err := client.Bucket("examplebucket")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
    err = bucket.PutObject("exampledir/exampleobject.txt", bytes.NewReader([]byte("yourObjectValueByteArrary")))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }
}