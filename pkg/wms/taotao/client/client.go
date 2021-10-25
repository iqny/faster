package client

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"orp/pkg/wms/interface_factory"
	"orp/pkg/wms/taotao/sign/hmac"
	"orp/pkg/wms/taotao/sign/hmac256"
	"orp/pkg/wms/taotao/sign/md5"
	"sort"
	"strings"
	"time"
)


var sysParams = map[string]string{
	"v":           "2.0",
	"format":      "xml",
	"sign_method": "md5",
	"method":      "",
	"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
}

type Client struct {
	c *interface_factory.Config
}

func (cli *Client) Execute(req interface_factory.Request) (res interface_factory.Response, err error) {
	sysParams["method"] = req.GetMethod()
	sysParams["app_key"] = cli.c.AppKey
	sysParams["customerId"] = cli.c.CustomerId
	if r, err := req.Check(); err != nil {
		return r, err
	}
	body := req.ToXML()
	sysParams["sign"] = cli.generateSign(body)
	return cli.post(body, cli.getUrl())

}
func (cli *Client) getUrl() string {
	return strings.Join([]string{cli.c.GatewayUrl, "?", cli.getHttpQuery()}, "")
}

// generateSign 生成sign
func (cli *Client) generateSign(body string) string {
	var sign string
	switch sysParams["sign_method"] {
	case "md5":
		sign = md5.GenerateMd5(cli.getSysParams(body))
	case "hmac":
		sign = hmac.GenerateMd5(strings.Join([]string{cli.getSysParams(body)}, cli.c.Secret), cli.c.Secret)
	case "hmac-sha256":
		sign = hmac256.GenerateHmacSha256(strings.Join([]string{cli.getSysParams(body)}, cli.c.Secret), cli.c.Secret)
	}
	/*w := md5.New()
	_, err := io.Writestring(w, cli.getSysParams(body))
	if err != nil {
		return ""
	}
	return strings.ToUpper(fmt.Sprintf("%x", w.Sum(nil)))*/
	return sign
}

// getSysParams 获取param
func (cli *Client) getSysParams(xmlBody string) string {
	var keys []string
	for k := range sysParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := cli.c.AppSecret
	for _, k := range keys {
		s += k + sysParams[k]
	}
	s += xmlBody + cli.c.AppSecret
	return s
}
func (cli *Client) getHttpQuery() string {
	sysParams["timestamp"] = url.QueryEscape(sysParams["timestamp"])
	var keys []string
	for k := range sysParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := "" //cli.c.appSecret
	for _, k := range keys {
		s += k + "=" + sysParams[k] + "&"
	}
	return strings.Trim(s, "&")
}
func (cli *Client) post(xmlBody, url string) (res interface_factory.Response, err error) {
	client := &http.Client{}
	var request *http.Request
	request, err = http.NewRequest("POST", url, strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(request.Body)
	var do *http.Response
	do, err = client.Do(request)
	if err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(do.Body)
	if err != nil {
		return
	}
	var succRes interface_factory.SuccessResponse
	err = xml.Unmarshal(body, &succRes)
	succRes.Res = string(body)
	succRes.Req = xmlBody
	res = succRes
	if err != nil {
		var errRes interface_factory.ErrResponse
		err = xml.Unmarshal(body, &errRes)
		errRes.Res = string(body)
		errRes.Req = xmlBody
		res = errRes
		if err != nil {
			return
		}
	}
	return
}
func New(cfs ...interface_factory.ConfigFunc) *Client {
	c := &interface_factory.Config{
		GatewayUrl: "http://qimen.api.taobao.com/router/qmtest",
		Secret:     "abc",
	}
	for _, f := range cfs {
		f(c)
	}
	return &Client{c: c}
}
