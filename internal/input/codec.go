package input

import "encoding/json"

func Codec(msg []byte, codec string) map[string]interface{} {
	field := map[string]interface{}{}
	err := json.Unmarshal(msg, &field)
	if err != nil {
		field["message"] = string(msg)
	}
	return field
}
