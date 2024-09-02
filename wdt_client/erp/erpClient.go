package erp

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
	nurl "net/url"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"
)

type (
	WdtClient struct {
		AppKey         string
		AppSecret      string
		Sid            string
		BaseUrl        string
		PageSize       string
		ConnectTimeout int
		ReadTimeout    int
	}
)

func NewWdtClient(clientParams map[string]string) (wdtClient *WdtClient) {
	return &WdtClient{
		AppKey:         clientParams["AppKey"],
		AppSecret:      clientParams["AppSecret"],
		Sid:            clientParams["Sid"],
		BaseUrl:        clientParams["BaseUrl"],
		PageSize:       "100",
		ConnectTimeout: 3000,
		ReadTimeout:    15000,
	}
}

func (c *WdtClient) signRequest(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	query := ""
	for _, k := range keys {
		if k == "sign" {
			continue
		}
		if query != "" {
			query += ";"
		}

		query += fmt.Sprintf("%02d-%s:", utf8.RuneCountInString(k), k)

		v := params[k]
		query += fmt.Sprintf("%04d-%s", utf8.RuneCountInString(v), v)
	}

	query += c.AppSecret

	harsher := md5.New()
	_, err := io.WriteString(harsher, query)
	if err != nil {
		return ""
	}
	sum := harsher.Sum(nil)
	md5str := hex.EncodeToString(sum)
	return md5str
}

func (c *WdtClient) Execute(method string, params map[string]string) (*map[string]interface{}, error) {
	params["appkey"] = c.AppKey
	params["sid"] = c.Sid
	params["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["sign"] = c.signRequest(params)

	req, _ := buildRequest(context.Background(), http.MethodPost, c.BaseUrl+method, params)
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

func (c *WdtClient) executeWithData(method string, params map[string]string, data any) (*map[string]interface{}, error) {
	params["appkey"] = c.AppKey
	params["sid"] = c.Sid
	params["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["sign"] = c.signRequest(params)

	err := mapstructure.Decode(params, &data)
	if err != nil {
		return nil, err
	}

	resp, err := httpc.Do(context.Background(), http.MethodPost, c.BaseUrl+method, data)
	if err != nil {
		return nil, err
	}

	respParse := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&respParse); err != nil {
		return nil, err
	}
	return &respParse, nil
}
