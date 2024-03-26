package convert

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

type ReflectTag struct {
	Value interface{}
}

func (c ReflectTag) FindKeys(tag string) []string {
	rType := reflect.TypeOf(c.Value).Elem()
	rValue := reflect.ValueOf(c.Value).Elem()
	count := rType.NumField()
	var partialFields []string
	for i := 0; i < count; i++ {
		// 获取 req 传入参数的 key 名称
		fKey := rType.Field(i).Name
		// 判断 key 对应 value的值类型，仅处理 value值类型为指针的 key
		// 注意：value值 不是指针类型的key，为Authorization 以及 MpOpenId，这2个参数，不参与部分更新字段的统计
		if rValue.FieldByName(fKey).Kind() == reflect.Ptr {
			// 指针 指向的值 不是nil时，添加到部分更新切片中。
			// 注意，Value类型的nil判断，只能有 IsNil()方法
			if !rValue.FieldByName(fKey).IsNil() {
				// 获取 json tag的名称
				fTag := strings.Split(rType.Field(i).Tag.Get(tag), ",")[0]
				partialFields = append(partialFields, fTag)
			}
		}
	}
	return partialFields
}

func (c ReflectTag) GetFilterData(tag string) string {
	reqValue := reflect.ValueOf(c.Value).Elem()
	reqType := reqValue.Type()
	filterData := make(map[string]string)
	for i := 0; i < reqType.NumField(); i++ {
		// 仅处理可选的 筛选参数
		if reqValue.Field(i).Kind() == reflect.Ptr {
			fTag := strings.Split(reqType.Field(i).Tag.Get(tag), ",")[0]
			if fTag == "" {
				continue
			}
			if reqValue.Field(i).Elem().Kind() == reflect.String {
				filterData[fTag] = reqValue.Field(i).Elem().String()
			} else if reqValue.Field(i).Elem().Kind() == reflect.Int32 || reqValue.Field(i).Elem().Kind() == reflect.Int64 {
				filterData[fTag] = strconv.FormatInt(reqValue.Field(i).Elem().Int(), 10)
			} else if reqValue.Field(i).Elem().Kind() == reflect.Bool {
				bJson, _ := json.Marshal(reqValue.Field(i).Elem().Bool())
				filterData[fTag] = string(bJson)
			}
		}
	}
	data, _ := json.Marshal(filterData)
	return string(data)
}

func (c ReflectTag) GetMap(tag string) map[string]interface{} {
	reqValue := reflect.ValueOf(c.Value).Elem()
	reqType := reqValue.Type()
	filterData := make(map[string]interface{})
	for i := 0; i < reqType.NumField(); i++ {
		// 仅处理可选的 筛选参数
		if reqValue.Field(i).Kind() == reflect.Ptr {
			fTag := strings.Split(reqType.Field(i).Tag.Get(tag), ",")[0]
			if fTag == "" {
				continue
			}
			if reqValue.Field(i).Elem().Kind() == reflect.String {
				filterData[fTag] = reqValue.Field(i).Elem().String()
			} else if reqValue.Field(i).Elem().Kind() == reflect.Int32 || reqValue.Field(i).Elem().Kind() == reflect.Int64 {
				filterData[fTag] = strconv.FormatInt(reqValue.Field(i).Elem().Int(), 10)
			} else if reqValue.Field(i).Elem().Kind() == reflect.Bool {
				bJson, _ := json.Marshal(reqValue.Field(i).Elem().Bool())
				filterData[fTag] = string(bJson)
			}
		}
	}
	return filterData
}
