package handler

import (
	"strconv"
)

func getParamInt(val string, defaultVal int) int {
	if len(val) == 0 {
		return defaultVal
	}
	iVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return iVal
}
