package wdt_qm

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
	nurl "net/url"
	"sort"
	"strings"
	"time"
)

type QmClient struct {
	Sid          string `form:"sid" mapstructure:"sid"`                       // wdt2
	AppKey       string `form:"app_key" mapstructure:"app_key"`               // 奇门授权应用（21xxxxxx）
	Timestamp    string `form:"timestamp" mapstructure:"timestamp"`           // 格式化后的时间戳2019-12-06 10:37:00
	Sign         string `form:"sign" mapstructure:"sign"`                     // 奇门签名计算规则
	TargetAppKey string `form:"target_app_key" mapstructure:"target_app_key"` // 旺店通在奇门授权应用（授权通知邮件中提供）
	Format       string `form:"format" mapstructure:"format"`                 // 固定值json
	SignMethod   string `form:"sign_method" mapstructure:"sign_method"`       // 固定值md5
	V            string `form:"v" mapstructure:"v"`                           // 固定值2.0
	BaseUrl      string `form:"base_url" mapstructure:"base_url"`             // 奇门接口地址

	AppSecret      string
	PageSize       string
	ConnectTimeout int
	ReadTimeout    int
}

// NewQmClient 旺店通 奇门客户端
func NewQmClient(clientParams map[string]string) *QmClient {
	return &QmClient{
		Sid:          clientParams["Sid"],
		AppKey:       clientParams["AppKey"],
		TargetAppKey: clientParams["TargetAppKey"],
		AppSecret:    clientParams["AppSecret"],
		BaseUrl:      clientParams["BaseUrl"],

		Format:     "json",
		SignMethod: "md5",
		V:          "2.0",

		PageSize:       "100",
		ConnectTimeout: 3000,
		ReadTimeout:    15000,
	}
}
func (c *QmClient) Execute(method string, params map[string]string) (*map[string]interface{}, error) {
	params["sid"] = c.Sid
	params["app_key"] = c.AppKey
	params["target_app_key"] = c.TargetAppKey
	params["method"] = method
	params["format"] = c.Format
	params["sign_method"] = c.SignMethod
	params["v"] = c.V
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["sign"] = c.signRequest(params)

	req, _ := buildRequest(context.Background(), http.MethodPost, c.BaseUrl, params)
	resp, err := httpc.DoRequest(req)
	if err != nil {
		return nil, err
	}

	respParse := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&respParse); err != nil {
		return nil, err
	}
	return &respParse, nil
}

func (c *QmClient) signRequest(params map[string]string) string {
	var allParams = make(map[string]interface{})

	for k, v := range params {
		allParams[k] = v
	}

	var keyList []string
	for k := range allParams {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	var signStr = ""
	for _, key := range keyList {
		value := allParams[key]
		signStr = signStr + fmt.Sprintf("%v%v", key, value)
	}

	signStr = c.AppSecret + signStr + c.AppSecret

	harsher := md5.New()
	_, err := io.WriteString(harsher, signStr)
	if err != nil {
		return ""
	}
	sum := harsher.Sum(nil)
	md5str := strings.ToUpper(hex.EncodeToString(sum))

	return md5str
}

func buildRequest(ctx context.Context, method, url string, params map[string]string) (*http.Request, error) {
	u, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}

	jsonValue, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bytes.NewReader(jsonValue))
	if err != nil {
		return nil, err
	}

	val := convertMapStringToStringInterface(params)
	req.URL.RawQuery = buildFormQuery(u, val)
	return req, nil
}

func buildFormQuery(u *nurl.URL, val map[string]any) string {
	query := u.Query()
	for k, v := range val {
		query.Add(k, fmt.Sprint(v))
	}

	return query.Encode()
}

func convertMapStringToStringInterface(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		result[k] = v
	}
	return result
}
