package utils

import "encoding/json"

func ParseToBytes(data interface{}) ([]byte, error) {
	dataBuffer, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return dataBuffer, nil
}
