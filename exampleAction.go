package simpleAliyunApiClient

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SendSms struct {
	PhoneNumbers    string `aliParam:"PhoneNumbers"`
	SignName        string `aliParam:"SignName"`
	TemplateCode    string `aliParam:"TemplateCode"`
	TemplateParam   string `aliParam:"TemplateParam"`
	SmsUpExtendCode string `aliParam:"SmsUpExtendCode"`
	OutId           string `aliParam:"OutId"`
}

func (action *SendSms) GetActionInfo() ActionInfo {
	return ActionInfo{
		Api:     "SendSms",
		Version: "2017-01-01",
	}
}

func (action *SendSms) SetPhoneNumbers(phoneNumbers []string) error {
	if len(phoneNumbers) > 1000 {
		return fmt.Errorf("批量发送短信，上限为1000个手机号码。")
	}

	action.PhoneNumbers = strings.Join(phoneNumbers, ",")
	return nil
}

func (action *SendSms) SetTemplateParam(param map[string]string) error {
	paramJsonStr, err := json.Marshal(param)
	if err != nil {
		return err
	}
	action.TemplateParam = string(paramJsonStr)
	return nil
}
