package model

import "time"

type Seckill struct {
	SeckillId  int64     `json:"seckillId"`
	Name       string    `json:"name"`
	Number     uint32    `json:"number"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	CreateTime time.Time `json:"createTime"`
}
