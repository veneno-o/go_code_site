package utils

import (
	"strconv"
	"time"
)

func GetCurTime() string {
	d := time.Now().UnixMilli() //毫秒
	s := strconv.FormatInt(d, 10)
	return s
}
