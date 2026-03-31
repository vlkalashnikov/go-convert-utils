package utils

import "reflect"

func ChunkArray[T any](arr []T, chunkSize int) (ret [][]T) {
	if chunkSize <= 0 {
		return nil
	}

	ret = make([][]T, 0, (len(arr)+chunkSize-1)/chunkSize)

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		ret = append(ret, arr[i:end])
	}
	return
}

// Deprecated: use ChunkArray.
func ChunckArray[T any](arr []T, chunkSize int) (ret [][]T) {
	return ChunkArray(arr, chunkSize)
}

func UniqueBy[T any, K comparable](s []T, keyFn func(obj T) K) []T {
	allKeys := make(map[K]struct{})
	result := make([]T, 0, len(s))
	for _, item := range s {
		key := keyFn(item)
		if _, exists := allKeys[key]; exists {
			continue
		}

		allKeys[key] = struct{}{}
		result = append(result, item)
	}
	return result
}

func UniqSlice[T any](s []T, distinctByKey func(obj T) interface{}) []T {
	comparableKeys := make(map[interface{}]struct{})
	nonComparableKeys := make([]interface{}, 0)
	result := make([]T, 0, len(s))

	for _, item := range s {
		key := distinctByKey(item)
		if isComparableKey(key) {
			if _, exists := comparableKeys[key]; exists {
				continue
			}
			comparableKeys[key] = struct{}{}
			result = append(result, item)
			continue
		}

		seen := false
		for _, existing := range nonComparableKeys {
			if reflect.DeepEqual(existing, key) {
				seen = true
				break
			}
		}
		if seen {
			continue
		}

		nonComparableKeys = append(nonComparableKeys, key)
		result = append(result, item)
	}

	return result
}

func isComparableKey(key interface{}) bool {
	if key == nil {
		return true
	}
	return reflect.TypeOf(key).Comparable()
}
