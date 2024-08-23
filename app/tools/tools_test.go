package tools

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	pwd := "user"
	EncryptV1(pwd)
	fmt.Println(EncryptV1(pwd))
}
