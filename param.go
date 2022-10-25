package simpleAliyunApiClient

import (
	"fmt"
	"reflect"
)

const tagName = "aliParam"

// StructToMap 将struct结构体转换为map
func StructToMap(structs ...any) (map[string]string, error) {
	result := make(map[string]string)

	for _, in := range structs {
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
			if tagValue := field.Tag.Get(TagName); tagValue != "" {
				result[tagValue] = fmt.Sprintf("%s", elem.Field(i).Interface())
			}
		}
	}

	return result, nil
}
