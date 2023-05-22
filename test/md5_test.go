package test

import (
	"fmt"
	"scaffold-gin/util"
	"testing"
)

func TestMd5(t *testing.T) {
	s := util.Md5String("123456")
	fmt.Println(s)
}