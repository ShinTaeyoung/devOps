package utils
import (
	"encoding/json"
	"strconv"
	"fmt"
)

type PayloadMap struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	TxId    string                 `json:"txId"`
}
type PayloadMapArray struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
	TxId    string                   `json:"txId"`
}

func GenerateResponseError(code int, message string) ([]byte) {
	payload := PayloadMap{Code: code, Message: message}
	response, err := json.Marshal(payload)
	if err != nil {
		return []byte(fmt.Sprint(503, strconv.Itoa(code), message))
	}
	return response;
}

func GenerateResponseBytes(code int, message string, data []byte, txId string) ([]byte) {
	
		var mapResponse map[string]interface{}
		json.Unmarshal(data, &mapResponse)
	
		payload := PayloadMap{Code: code, Message: message, Data: mapResponse, TxId: txId}
	
		response, err := json.Marshal(payload)
		if err != nil {
			return []byte(fmt.Sprint(503, strconv.Itoa(code), message))
		}
		return response;
	}