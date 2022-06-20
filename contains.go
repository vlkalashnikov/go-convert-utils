package utils

func SInArr(values []string, value string) bool {
	for _, elem := range values {
		if elem == value {
			return true
		}
	}
	return false
}

func IInArr(ids []int, id int) bool {
	for _, d := range ids {
		if id == d {
			return true
		}
	}
	return false
}
