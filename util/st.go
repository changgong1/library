package util

import "encoding/json"

func StrToSt(str string, ret interface{}) error {
	return json.Unmarshal([]byte(str), ret)
}

func StToStr(data interface{}) string {
	strBody, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(strBody)
}

func StToMap(data interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	str := StToStr(data)
	_ = json.Unmarshal([]byte(str), &ret)
	return ret
}

func MapToSt(data interface{}, st interface{}) error {
	body := StToStr(data)
	return StrToSt(body, st)
}
