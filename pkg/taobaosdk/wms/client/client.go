package client

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"orp/pkg/taobaosdk/wms/request"
	"orp/pkg/taobaosdk/wms/response"
	"orp/pkg/taobaosdk/wms/sign/hmac"
	"orp/pkg/taobaosdk/wms/sign/hmac256"
	"orp/pkg/taobaosdk/wms/sign/md5"
	"sort"
	"strings"
	"time"
)

type Config struct {
	appKey     string
	appSecret  string
	Session    string
	customerId string
	gatewayUrl string
	secret     string //用于hmac加密
}
type ConfigFunc func(c *Config)

var sysParams = map[string]string{
	"v":           "2.0",
	"format":      "xml",
	"sign_method": "md5",
	"method":      "",
	"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
}

type client struct {
	c *Config
}

func (cli *client) Execute(req request.Request) (res response.Response, err error) {
	sysParams["method"] = req.GetMethod()
	sysParams["app_key"] = cli.c.appKey
	sysParams["customerId"] = cli.c.customerId
	if r, err := req.Check(); err != nil {
		return r, err
	}
	body := req.ToXML()
	sysParams["sign"] = cli.generateSign(body)
	return cli.post(body, cli.getUrl())

}
func (cli *client) getUrl() string {
	return strings.Join([]string{cli.c.gatewayUrl, "?", cli.getHttpQuery()}, "")
}

// generateSign 生成sign
func (cli *client) generateSign(body string) string {
	var sign string
	switch sysParams["sign_method"] {
	case "md5":
		sign = md5.GenerateMd5(cli.getSysParams(body))
	case "hmac":
		sign = hmac.GenerateMd5(strings.Join([]string{cli.getSysParams(body)}, cli.c.secret), cli.c.secret)
	case "hmac-sha256":
		sign = hmac256.GenerateHmacSha256(strings.Join([]string{cli.getSysParams(body)}, cli.c.secret), cli.c.secret)
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
func (cli *client) getSysParams(xmlBody string) string {
	var keys []string
	for k := range sysParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := cli.c.appSecret
	for _, k := range keys {
		s += k + sysParams[k]
	}
	s += xmlBody + cli.c.appSecret
	return s
}
func (cli *client) getHttpQuery() string {
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
func (cli *client) post(xmlBody, url string) (res response.Response, err error) {
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
	var succRes response.SuccessResponse
	err = xml.Unmarshal(body, &succRes)
	succRes.Res = string(body)
	succRes.Req = xmlBody
	res = succRes
	if err != nil {
		var errRes response.ErrResponse
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
func New(cfs ...ConfigFunc) *client {
	c := &Config{
		gatewayUrl: "http://qimen.api.taobao.com/router/qmtest",
		secret:     "abc",
	}
	for _, f := range cfs {
		f(c)
	}
	return &client{c: c}
}
