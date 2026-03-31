package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	if a == nil {
		return errors.New("scan on nil JsonB pointer")
	}

	var b []byte
	switch v := value.(type) {
	case nil:
		*a = nil
		return nil
	case []byte:
		if len(v) == 0 {
			*a = JsonB{}
			return nil
		}
		b = v
	case string:
		if v == "" {
			*a = JsonB{}
			return nil
		}
		b = []byte(v)
	default:
		return fmt.Errorf("unsupported type %T", value)
	}

	var decoded JsonB
	if err := json.Unmarshal(b, &decoded); err != nil {
		return err
	}

	*a = decoded
	return nil
}
