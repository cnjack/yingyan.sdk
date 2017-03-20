package yingyan

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type s struct {
	ak  string
	sk  string
	serviceID int
	s   bool
}

// visit http://lbsyun.baidu.com/apiconsole/key/ 获取ak
// 如果设置的白名单则设置sk ""
func NewClient(ak, sk string,serviceID int) *s {
	return &s{
		sk: sk,
		ak: ak,
		s: true,
		serviceID:serviceID,
	}
}

func (serv *s) Post(path string, param map[string]string) (body []byte, err error) {
	data := url.Values{}
	data.Add("ak", serv.ak)
	data.Add("service_id", strconv.Itoa(serv.serviceID))
	for k, v := range param {
		data.Add(k, v)
	}
	sn := serv.sign(path, &data)
	if sn != "" {
		data.Add("sn", sn)
	}
	resp, err := http.PostForm(apiRootPath+path, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}