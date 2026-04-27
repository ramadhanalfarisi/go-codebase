package helpers

import "strconv"

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		Error(err)
		return 0
	}
	return result
}