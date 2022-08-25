package utils

func GetMapValue(mp map[string]interface{}, name string) interface{} {
	if x, found := mp[name]; found {
		return x
	}
	return nil
}

func GetIntValue(mp map[string]interface{}, name string) int {
	x := GetMapValue(mp, name)

	if val, ok := x.(float64); ok {
		return int(val)
	}

	if val, ok := x.(float32); ok {
		return int(val)
	}

	if val, ok := x.(int32); ok {
		return int(val)
	}

	if val, ok := x.(int64); ok {
		return int(val)
	}
	return 0
}

func GetStringValue(mp map[string]interface{}, name string) string {
	x := GetMapValue(mp, name)
	if val, ok := x.(string); ok {
		return val
	}
	return ""
}
