package tools

import "encoding/json"

func ToJson(obj interface{}) string {
	return string(ToJsonBs(obj))
}

func ToJsonBs(obj interface{}) []byte {
	bs, err := json.Marshal(obj)
	if err != nil {
		return []byte{}
	}
	return bs
}
