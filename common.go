package simpleAliyunApiClient

const TagName = "aliParam"

type AliApiAction interface {
	GetActionInfo() ActionInfo
}

type ActionInfo struct {
	Api     string
	Version string
}

type client struct {
	Format           string `aliParam:"Format"`           // 指定接口返回数据的格式。可以选择 json 或者 XML。默认为 XML
	AccessKeyId      string `aliParam:"AccessKeyId"`      // 阿里云访问密钥ID
	AccessSecret     string ``                            // 阿里云访问秘钥
	SignatureNonce   string `aliParam:"SignatureNonce"`   // 签名唯一随机数
	Timestamp        string `aliParam:"Timestamp"`        // 请求的时间戳。按照ISO8601标准表示，并需要使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ
	SignatureMethod  string `aliParam:"SignatureMethod"`  // 签名方式。目前为固定值 HMAC-SHA1
	SignatureVersion string `aliParam:"SignatureVersion"` // 签名算法版本。目前为固定值 1.0
	// Signature        string ``                            // 请求签名，用户请求的身份验证。
}
