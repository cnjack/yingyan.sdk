package yingyan

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

func (serv *s) sign(uri string, param *url.Values) (sn string) {
	if serv.sk == "" {
		return
	}
	o := uri + "?" + param.Encode() + serv.sk
	hash := md5.New()
	hash.Write([]byte(url.QueryEscape(o)))
	sn = hex.EncodeToString(hash.Sum(nil))
	return
}
