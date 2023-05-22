package test

import (
	"fmt"
	"regexp"
	"testing"
)


func TestReg(t *testing.T) {
	matched, err := regexp.MatchString("^1(3|4|5|6|8|9)[0-9]{9}$", "1308765051")
    fmt.Println(matched,err)
}