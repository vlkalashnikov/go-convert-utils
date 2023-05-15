package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type JsonB map[string]interface{} //@Name JSONB

// Value Marshal
func (a JsonB) Value() (driver.Value, error) {
	result := JsonB{}
	for k, v := range a {
		result[strings.ToLower(k)] = v
	}
	return json.Marshal(result)
}

// Scan Unmarshal
func (a *JsonB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
