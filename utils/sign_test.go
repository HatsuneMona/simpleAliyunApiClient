package utils

import (
	"testing"
	"time"
)

func Test_getRandomNonce(t *testing.T) {
	t.Run("test 5 times", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(1) * 321 * time.Nanosecond)
			got := GetRandomNonce(i * 5)
			t.Logf("%v time, getRandomNonce(lenth:%v) result: %s", i, i*5, got)
		}
	})
}

func Test_getUTCTimestamp(t *testing.T) {
	t.Run("test 5 times", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			got := GetUTCTimestamp()
			t.Logf("%v time, getUTCTimestamp() result: %s", i, got)
			time.Sleep(time.Duration(1) * time.Second)
		}
	})
}

func Test_percentEncode(t *testing.T) {
	// testStr := map[string]string{
	// 	"不编码字符": "abc123DEF-_._~",
	// 	"其它字符编码成%XY的格式，其中XY是字符对应ASCII码的16进制": `""{}[]:;<>,/?=!@#$%^&` + "`",
	// 	"中文": "你好",
	// 	"扩展的UTF-8字符 emoji 【😊    😘    😆】": "😊    😘    😆",
	// 	"空格【 】（预期【%20】）":                 " ",
	// 	"加号【+】（预期【%20】）":                 "+",
	// 	"星号【*】（预期【%2A】）":                 "*",
	// 	"波浪【~】（预期【~】）":                   "~",
	// }
	// for msg, str := range testStr {
	// 	fmt.Printf("测试项 %s 的结果为【%s】\n", msg, url.QueryEscape(str))
	// }

	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"不编码字符",
			args{"abc123DEF-_._~"},
			"abc123DEF-_._~",
		},
		{
			"其它字符编码成%XY的格式，其中XY是字符对应ASCII码的16进制",
			args{`""{}[]:;<>,/?=!@#$%^&` + "`"},
			"%22%22%7B%7D%5B%5D%3A%3B%3C%3E%2C%2F%3F%3D%21%40%23%24%25%5E%26%60",
		},
		{
			"中文",
			args{"你好"},
			"%E4%BD%A0%E5%A5%BD",
		},
		{
			"扩展的UTF-8字符 emoji 【😆】",
			args{"😆"},
			"%F0%9F%98%86",
		},
		{
			"空格【 】",
			args{" "},
			"%20",
		},
		{
			"加号【+】",
			args{"+"},
			"%20",
		},
		{
			"星号【*】",
			args{"*"},
			"%2A",
		},
		{
			"波浪【~】",
			args{"~"},
			"~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentEncode(tt.args.src); got != tt.want {
				t.Errorf("percentEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildSignStr(t *testing.T) {
	type args struct {
		paramMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty args",
			args{map[string]string{}},
			"GET&%2F",
		},
		{
			"sort args",
			args{map[string]string{
				"C": "apple",
				"A": "apple",
				"B": "apple",
			}},
			"GET&%2F&A=apple&B=apple&C=apple",
		},
		{
			"empty key or value",
			args{map[string]string{
				"C": "apple",
				"":  "apple",
				"B": "",
			}},
			"GET&%2F&C=apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildSignStr(tt.args.paramMap); got != tt.want {
				t.Errorf("BuildSignStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
