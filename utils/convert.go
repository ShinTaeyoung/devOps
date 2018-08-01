package utils

import (
	"encoding/json"
)

func ArrayToMap(v interface{}) ([]map[string]interface{}) {
	data, _ := json.Marshal(v)
	var mapResponse []map[string]interface{}
	json.Unmarshal(data, &mapResponse)
	return mapResponse
}