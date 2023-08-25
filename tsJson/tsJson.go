package tsJson

import (
	json2 "encoding/json"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToByte(obj interface{}) []byte {
	str, err := json.Marshal(obj)
	if err != nil {
		return make([]byte, 0)
	}
	return str
}

// ToJson obj不能为map类型
func ToJson(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(str)
}

func ToString(obj interface{}) (string, error) {
	str, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func FromString(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), obj)
	return err
}

// 解决原始toJson写法无法解析map问题
func ToJsonOriginal(obj interface{}) string {
	str, err := json2.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(str)
}

func ToByteOriginal(obj interface{}) []byte {
	str, err := json2.Marshal(obj)
	if err != nil {
		return []byte{}
	}
	return str
}

func FromStringOriginal(str string, obj interface{}) error {
	err := json2.Unmarshal([]byte(str), obj)
	return err
}
