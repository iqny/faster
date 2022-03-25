package top_client

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"orp/pkg/sdk/topclient/sign/hmac"
	"orp/pkg/sdk/topclient/sign/hmac256"
	"orp/pkg/sdk/topclient/sign/md5"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"
)

type Request interface {
	GetApiMethodName() string
	GetApiParas() map[string]string
}
type Response interface {
	GetCode() string
	GetContent() string
}
type Config struct {
	appKey     string
	appSecret  string
	session    string
	gatewayUrl string
	secret     string //用于hmac加密
	format     string
	signMethod string
	v          string
}
type ConfigFunc func(c *Config)

// SuccessResponse x := "<response><flag>success</flag><code>200</code><message>Success</message><data></data></response>"
type SuccessResponse struct {
	XMLName    xml.Name    `xml:"response"`
	Flag       string      `xml:"flag"`
	Code       string      `xml:"code"`
	Message    string      `xml:"message"`
	SubCode    string      `xml:"sub_code"`
	SubMessage string      `xml:"sub_message"`
	RequestId  string      `xml:"request_id"`
	Data       interface{} `xml:"data"`
	Res        string      `xml:"-"`
	Req        string      `xml:"-"`
}
type SuccessResponseJson map[string]interface{}

func (r SuccessResponseJson) GetCode() string {
	return ""
}
func (r SuccessResponseJson) GetContent() string {
	return ""
}
func (r SuccessResponse) GetCode() string {
	return r.Code
}
func (r SuccessResponse) GetContent() string {
	return r.Res
}

// ErrResponse xres:="<error_response><code>50</code><msg>Remote service error</msg><sub_code>isv.invalid-parameter</sub_code><sub_message>非法参数</sub_message></error_response>"
type ErrResponse struct {
	XMLName    xml.Name `xml:"error_response"`
	Code       string   `xml:"code"`
	Message    string   `xml:"msg"`
	SubCode    string   `xml:"sub_code"`
	SubMessage string   `xml:"sub_message"`
	RequestId  string   `xml:"request_id"`
	Res        string   `xml:"-"`
	Req        string   `xml:"-"`
}

func (r ErrResponse) GetCode() string {
	return r.Code
}
func (r ErrResponse) GetContent() string {
	return r.Res
}

type Client struct {
	c         *Config
	sysParams map[string]string
	logHandel *os.File
}
func New(cfs ...ConfigFunc) *Client {
	c := &Config{
		gatewayUrl: "https://gw.api.taobao.com/router/rest",
		secret:     "abc",
		v:          "2.0",
		format:     "xml",
		signMethod: "md5",
	}
	for _, f := range cfs {
		f(c)
	}
	return &Client{
		c: c,
		sysParams: map[string]string{
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
}
func (cli *Client) Execute(req Request, session string) (res SuccessResponseJson, err error) {
	cli.sysParams["method"] = req.GetApiMethodName()
	cli.sysParams["app_key"] = cli.c.appKey
	cli.sysParams["sign_method"] = cli.c.signMethod
	cli.sysParams["v"] = cli.c.v
	cli.sysParams["format"] = cli.c.format
	if session != "" {
		cli.sysParams["session"] = session
	}
	apiParams := req.GetApiParas()
	cli.sysParams["sign"] = cli.generateSign(merge(apiParams, cli.sysParams))
	response, err := cli.post(cli.getUrl(), apiParams)
	if err == nil {
		switch response.(type) {
		case SuccessResponseJson:
			res = response.(SuccessResponseJson)
		}
	}
	return

}
func (cli *Client) getUrl() string {
	return strings.Join([]string{cli.c.gatewayUrl, "?", cli.getHttpQuery()}, "")
}

// generateSign 生成sign
func (cli *Client) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := cli.c.appSecret
	for _, k := range keys {
		s += k + params[k]
	}
	s += cli.c.appSecret
	var sign string
	switch cli.sysParams["sign_method"] {
	case "md5":
		sign = md5.GenerateMd5(s)
	case "hmac":
		sign = hmac.GenerateMd5(strings.Join([]string{s}, cli.c.secret), cli.c.secret)
	case "hmac-sha256":
		sign = hmac256.GenerateHmacSha256(strings.Join([]string{s}, cli.c.secret), cli.c.secret)
	}
	/*w := md5.New()
	_, err := io.Writestring(w, cli.getSysParams(body))
	if err != nil {
		return ""
	}
	return strings.ToUpper(fmt.Sprintf("%x", w.Sum(nil)))*/
	return sign
}

func (cli *Client) getHttpQuery() string {
	var keys []string
	for k := range cli.sysParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := "" //cli.c.appSecret
	for _, k := range keys {
		s += k + "=" + url.QueryEscape(cli.sysParams[k]) + "&"
	}
	return strings.Trim(s, "&")
}
func (cli *Client) post(uri string, params map[string]string) (res Response, err error) {
	s := ""
	for k, v := range params {
		s += k + "=" + url.QueryEscape(v) + "&"
	}
	s = strings.Trim(s, "&")
	client := &http.Client{}
	var request *http.Request
	request, err = http.NewRequest("POST", uri, strings.NewReader(s))
	if err != nil {
		cli.log("POST:", err)
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			cli.log("close:", err)
			return
		}
	}(request.Body)
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		cli.log("client.Do:", err)
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)

	if err != nil {
		cli.log("reader:", err)
		return
	}
	if cli.c.format == "xml" {
		var succRes SuccessResponse
		err = xml.Unmarshal(body, &succRes)
		succRes.Res = string(body)
		//succRes.Req = xmlBody
		res = succRes
		if err != nil {
			var errRes ErrResponse
			err = xml.Unmarshal(body, &errRes)
			errRes.Res = string(body)
			//errRes.Req = xmlBody
			res = errRes
			if err != nil {
				return
			}
		}
	} else {
		var succRes, subResponse SuccessResponseJson
		err = json.Unmarshal(body, &succRes)
		subResponse = make(map[string]interface{})
		for _, sub := range succRes {
			for k, v := range sub.(map[string]interface{}) {
				subResponse[k] = v
			}
			res = subResponse
		}

	}
	return
}

func merge(array1, arrar2 map[string]string) map[string]string {
	newArrar := make(map[string]string, len(array1)+len(arrar2))
	for i, v := range array1 {
		newArrar[i] = v
	}
	for i, v := range arrar2 {
		newArrar[i] = v
	}
	return newArrar
}
func (cli *Client) log(msg ...interface{}) {
	//var err error
	var tempLogDir, temp string
	//如果
	if runtime.GOOS == "windows" {
		temp = os.Getenv("TEMP")
	}
	if temp == "" {
		temp = "/tmp/"
	}
	tempLogDir = temp + "/top_client/logs/"
	if _, err := os.Stat(tempLogDir); os.IsNotExist(err) {
		os.MkdirAll(tempLogDir, 0766)
	}
	logFile, _ := os.OpenFile(tempLogDir+cli.c.appKey+"-"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	log.SetOutput(logFile)
	log.Println(msg...)
}
