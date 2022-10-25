package simpleAliyunApiClient

import (
	"fmt"
)

func New(ak, as string) (*client, error) {
	if ak == "" || as == "" {
		return nil, fmt.Errorf("accessKeyId or accessSecret empty")
	}

	// 初始化公共参数
	return &client{
		Format:           "json",
		AccessKeyId:      ak,
		AccessSecret:     as,
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
	}, nil
}

func (c *client) Do(action AliApiAction) error {
	api := action.GetActionInfo()
	fmt.Printf("Do! action:%+v ;client:%+v \n", api.Api, api.Version)

	param, err := StructToMap(action, c)
	if err != nil {
		return err
	}

	fmt.Printf("param map: %+v", param)

	return nil
}
