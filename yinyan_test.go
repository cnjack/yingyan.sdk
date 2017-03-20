package yingyan_test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	u, _ := url.Parse("http://baidu.com/geocoder/v2/?ak=xxxx")
	fmt.Println(u.Scheme + "://" + u.Host + u.EscapedPath())
	u.RequestURI()
}
