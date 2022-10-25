package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"time"
)

// GetRandomNonce 获取一个长度为 length 的随机字符串
func GetRandomNonce(length int) string {
	base := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		index := r.Intn(len(base))
		result[i] = base[index]
	}
	return string(result)
}

// GetUTCTimestamp 获取当前UTC时间戳
func GetUTCTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

// BuildSignStr 构造签名所需的字符串
func BuildSignStr(paramMap map[string]string) string {
	sortKey := make([]string, 0, len(paramMap))
	for key, value := range paramMap {
		if key != "" && value != "" {
			sortKey = append(sortKey, key)
		}
	}

	sort.Strings(sortKey)

	signStr := "GET&" + percentEncode("/")

	for _, key := range sortKey {
		signStr = fmt.Sprint(signStr, "&", key, "=", percentEncode(paramMap[key]))
	}

	return signStr
}

// percentEncode 参数编码
// 在阿里云 OpenAPI 调用中，我们需要对请求参数和请求值，使用 UTF-8 字符集按照RFC3986规则进行编码。具体编码规则如下：
// - 字符 A~Z、a~z、0~9 以及字符-、_、.、~不编码。
// - 对其它 ASCII 码字符进行编码。编码格式为%加上16进制的 ASCII 码。例如半角双引号（"）将被编码为 %22。
// - 非 ASCII 码通过 UTF-8 编码。
// - 空格编码成%20，而不是加号（+）。
func percentEncode(src string) string {
	t := []rune(src)
	str := make([]rune, 0, len(t))

	for _, char := range t {
		switch char {
		case '+', ' ':
			str = append(str, []rune(" 20")...)
		default:
			str = append(str, char)
		}
	}

	dst := url.QueryEscape(string(str))
	return strings.ReplaceAll(dst, "+", "%")
}
