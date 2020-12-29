package util

import (
	"log"
	"strconv"
	"time"
)

func ToTime(s string) (time.Time, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Printf("time [%s] string can not convert int", s)
		return time.Now(), err
	}

	return time.Unix(int64(i), 0), nil
}
