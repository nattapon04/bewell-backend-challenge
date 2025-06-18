package appjson

import (
	"encoding/json"
)

func Stringtify(jsons interface{}) (string, error) {
	metaByte, err := json.Marshal(jsons)

	return string(metaByte), err
}

func Parse(jsons interface{}, data string) error {
	err := json.Unmarshal([]byte(data), &jsons)

	return err
}

func ParseTo(data interface{}, toData interface{}) error {
	metaByte, err := Stringtify(data)
	if err != nil {
		return err
	}

	err = Parse(toData, metaByte)
	if err != nil {
		return err
	}

	return nil
}

func KeyValueToJSON(fields ...interface{}) (string, error) {
	data := make(map[string]interface{})

	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue // or return an error if key isn't a string
		}
		data[key] = fields[i+1]
	}

	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
