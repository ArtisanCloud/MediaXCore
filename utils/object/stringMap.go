package object

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type StringMap map[string]string

// ------------------------------- Merge --------------------------------------------
func MergeStringMap(toStringMap *StringMap, subStringMaps ...*StringMap) *StringMap {
	if toStringMap == nil {
		toStringMap = &StringMap{}
	}
	for _, subStringMap := range subStringMaps {
		if subStringMap != nil {
			// 迭代每个HashMap
			for k, v := range *subStringMap {
				toV := (*toStringMap)[k]
				// if the key is not exist in toMap
				if toV == "" && v != "" {
					(*toStringMap)[k] = v
				}

			}
		}
	}
	return toStringMap
}

// ------------------------------- Replace --------------------------------------------
func ReplaceStringMapRecursive(toMap *StringMap, subMaps ...*StringMap) *StringMap {
	if toMap == nil {
		toMap = &StringMap{}
	}
	// 拍平subMaps
	for _, subMap := range subMaps {
		if subMap != nil {
			// 迭代每个HashMap
			for k, v := range *subMap {
				(*toMap)[k] = v
			}
		}
	}
	return toMap
}

func ConvertStringMapToString(m *StringMap, separate string) string {
	var b bytes.Buffer
	for key, value := range *m {
		fmt.Fprintf(&b, "%s=%s%s", key, value, separate)
	}
	//fmt.Fprint(&b, "/0")
	return b.String()
}

// ------------------------------- Conversion ---------------------------------------

func StructToStringMapWithTag(obj interface{}, tag string) (newMap *StringMap, err error) {

	newMap = &StringMap{}

	if obj == nil {
		return newMap, err
	}

	e := reflect.ValueOf(obj).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i).Interface()
		var strField string = ""

		strField = fmt.Sprintf("%v", field)
		key := e.Type().Field(i).Name
		if tag != "" {
			key = e.Type().Field(i).Tag.Get(tag)
		}
		(*newMap)[key] = strField

	}

	return newMap, err

}

//func StructToStringMap(obj interface{}) (newMap *StringMap, err error) {
//	data, err := json.Marshal(obj) // Convert to a json string
//
//	if err != nil {
//		return
//	}
//
//	newMap = &StringMap{}
//	err = json.Unmarshal(data, newMap) // Convert to a map
//	return
//}

func StructToStringMap(obj interface{}) (*StringMap, error) {
	// 使用反射遍历字段
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct type but got %s", val.Kind())
	}

	newMap := make(StringMap) // 假设 StringMap = map[string]string

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 获取字段名或 tag
		key := fieldType.Tag.Get("json")
		if key == "-" {
			continue // 跳过 json:"-" 的字段
		}
		if key == "" {
			key = fieldType.Name
		}
		// 跳过未导出字段
		if !field.CanInterface() {
			continue
		}

		// 转换值为字符串
		valueStr := fmt.Sprintf("%v", field.Interface())
		newMap[key] = valueStr
	}

	return &newMap, nil
}

func GetJoinedWithKSort(params *StringMap) string {

	var strJoined string

	// ksort
	var keys []string
	for k := range *params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// join
	for _, k := range keys {
		strJoined += k + "=" + (*params)[k] + "&"
	}

	strJoined = strJoined[0 : len(strJoined)-1]

	return strJoined
}

// ------------------------------- Search --------------------------------------------

func InStringMap(val string, maps *StringMap) (exists bool, keys []string) {
	exists = false
	mapLen := len(*maps)
	keys = make([]string, 0, mapLen)

	for k, v := range *maps {
		if v == val {
			keys = append(keys, k)

			exists = true
		}
	}

	return exists, keys
}

func GetStringMapKV(maps StringMap) (keys []string, values []string) {
	mapLen := len(maps)
	keys = make([]string, 0, mapLen)
	values = make([]string, 0, mapLen)

	for k, v := range maps {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// ------------------------------- Filter --------------------------------------------
func FilterEmptyStringMap(mapData *StringMap) (filteredMap *StringMap) {
	filteredMap = &StringMap{}
	for k, v := range *mapData {
		if v != "" {
			(*filteredMap)[k] = v
		}
	}
	return filteredMap
}
