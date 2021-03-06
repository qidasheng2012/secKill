package util

import (
	"log"
	"strconv"
	"time"
)

func ToTime(s string) (time.Time, error) {
	if len(s) > 10 {
		s = string([]rune(s)[:10])
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		log.Printf("time [%s] can't convert int", s)
		return time.Now(), err
	}

	return time.Unix(int64(i), 0), nil
}
