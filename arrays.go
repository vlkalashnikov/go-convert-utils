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
