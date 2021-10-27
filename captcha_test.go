package goo_captcha

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	rsp := Get(80, 240, 4)
	fmt.Println(rsp)
}
