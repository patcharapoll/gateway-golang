package convert

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httputil"
)

// ConvertToReader ...
func ConvertToReader(data interface{}) (*bytes.Reader, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(dataBytes), nil
}

// ConvertJSONByteToStruct ...
func ConvertJSONByteToStruct(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}

//ConvertStructToJSONByte ...
func ConvertStructToJSONByte(data interface{}) ([]byte, error) {
	dataBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

//ConvertRequestToString ...
func ConvertRequestToString(req *http.Request) (string, error) {
	data, err := httputil.DumpRequest(req, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

//ConvertResponseToString ...
func ConvertResponseToString(resp *http.Response) (string, error) {
	data, err := httputil.DumpResponse(resp, true)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ConvertStructToStruct ...
func ConvertStructToStruct(input interface{}, target interface{}) (err error) {
	jsonByte, err := json.Marshal(input)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, target)

	if err != nil {
		return err
	}

	return
}
