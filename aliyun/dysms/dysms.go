package dysms

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DysmsClient struct {
	Client        *dysmsapi.Client
	SignName      string
	TemplateCode  string
	TemplateParam string
}

type AliyunSmsConf struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
}

func NweDysmsClient(conf *AliyunSmsConf) *DysmsClient {
	config := &openapi.Config{
		AccessKeyId:     tea.String(conf.AccessKeyId),
		AccessKeySecret: tea.String(conf.AccessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	client, _ := dysmsapi.NewClient(config)
	return &DysmsClient{
		Client: client,
	}
}

func (l *DysmsClient) SendSms(req *dysmsapi.SendSmsRequest, param *TemplateParamReq) (resp *DysmsResponse) {
	if param != nil {
		buff, _ := json.Marshal(param)
		req.TemplateParam = tea.String(string(buff))
	}
	result, err := l.Client.SendSms(req)
	resp = new(DysmsResponse)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
		return
	}
	if !tea.BoolValue(util.EqualString(result.Body.Code, tea.String("OK"))) {
		resp.Code = -1
		resp.Msg = tea.StringValue(result.Body.Message)
		return
	}
	resp.Msg = "OK"
	return
}

func (l *DysmsClient) SendServiceStatus(phone string, param TemplateParamReq) (resp *DysmsResponse) {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers: &phone,
		SignName:     tea.String(SignNameServer),
		TemplateCode: tea.String(TemplateCode_Care_ReviewResult),
	}
	return l.SendSms(req, &param)
}

func (l *DysmsClient) SendDeliver(phone string, param TemplateParamReq) (resp *DysmsResponse) {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers: &phone,
		SignName:     tea.String(SignNameServer),
		TemplateCode: tea.String(TemplateCode_Care_DeliverGoods),
	}
	return l.SendSms(req, &param)
}
