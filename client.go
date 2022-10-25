package simpleAliyunApiClient

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"reflect"
	"simpleAliyunApiClient/utils"
)

const tagName = "aliParam"

func New(ak, as string) (*Client, error) {
	if ak == "" || as == "" {
		return nil, fmt.Errorf("accessKeyId or accessSecret empty")
	}

	// 初始化公共参数
	return &Client{
		Format:           "json",
		AccessKeyId:      ak,
		accessSecret:     as,
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
	}, nil
}

func (c *Client) Do(action AliApiAction) error {
	api := action.GetActionInfo()
	fmt.Printf("Do! action:%+v ;Client:%+v \n", api.Api, api.Version)

	c.Timestamp = utils.GetUTCTimestamp()
	c.SignatureNonce = utils.GetRandomNonce(16)
	defer func() {
		c.Timestamp = ""
		c.SignatureNonce = ""
	}()

	signature, err := c.signRequest(action)
	if err != nil {
		return err
	}
	c.Signature = signature

	return nil
}

// signRequest 计算签名
func (c *Client) signRequest(action AliApiAction) (string, error) {
	param, err := c.buildParamMap(action)
	if err != nil {
		return "", err
	}
	fmt.Printf("param map: %+v", param)

	signSourceStr := utils.BuildSignStr(param)
	mac := hmac.New(sha1.New, []byte(c.accessSecret+"&"))
	mac.Write([]byte(signSourceStr))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// buildParamMap 将所有 Aliyun API 必备的 参数 转换为map
func (c *Client) buildParamMap(action AliApiAction) (map[string]string, error) {
	result := make(map[string]string)

	for _, in := range []any{c, action} {
		elem := reflect.ValueOf(in)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		if elem.Kind() != reflect.Struct { // 非结构体返回错误提示
			return nil, fmt.Errorf("action only accepts struct or struct pointer; got %T", elem)
		}

		// fmt.Printf("elem: %+v \n", elem)

		elemType := elem.Type()
		for i := 0; i < elemType.NumField(); i++ {
			// 获取每个成员的结构体字段类型
			field := elemType.Field(i)
			// 输出成员名和tag
			// fmt.Printf("name: %v  tag: '%v' tagValue: ", field.Name, field.Tag)
			if tagValue := field.Tag.Get(tagName); tagValue != "" {
				result[tagValue] = fmt.Sprintf("%s", elem.Field(i).Interface())
			}
		}
	}

	return result, nil
}
