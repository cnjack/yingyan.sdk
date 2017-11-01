package yingyan

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	ak         string
	sk         string
	httpClient *http.Client
	serviceID  int
	s          bool
}

// visit http://lbsyun.baidu.com/apiconsole/key/ 获取ak
// 如果设置的白名单则设置sk ""
func NewClient(ak, sk string, serviceID int) *Client {
	return &Client{
		sk:         sk,
		ak:         ak,
		s:          true,
		serviceID:  serviceID,
		httpClient: &http.Client{},
	}
}

// SetHttpClient you can set your own http client
func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) Post(path string, param map[string]string) (body []byte, err error) {
	data := url.Values{}
	data.Add("ak", c.ak)
	data.Add("service_id", strconv.Itoa(c.serviceID))
	for k, v := range param {
		data.Add(k, v)
	}
	sn := c.sign(path, &data)
	if sn != "" {
		data.Add("sn", sn)
	}
	resp, err := c.httpClient.PostForm(apiRootPath+path, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) Get(path string, param map[string]string) (body []byte, err error) {
	data := url.Values{}
	data.Add("ak", c.ak)
	data.Add("service_id", strconv.Itoa(c.serviceID))
	for k, v := range param {
		data.Add(k, v)
	}
	sn := c.sign(path, &data)
	if sn != "" {
		data.Add("sn", sn)
	}
	resp, err := c.httpClient.Get(apiRootPath + path + "?" + data.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
