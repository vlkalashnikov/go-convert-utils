package utils

func ChunckArray[T any](arr []T, chunkSize int) (ret [][]T) {
	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		ret = append(ret, arr[i:end])
	}
	return
}

func UniqSlice[T any](s []T, distinctByKey func(obj T) interface{}) []T {
	allKeys := make(map[interface{}]bool)
	var result []T
	for _, item := range s {
		if _, value := allKeys[distinctByKey(item)]; !value {
			allKeys[distinctByKey(item)] = true
			result = append(result, item)
		}
	}
	return result
}
