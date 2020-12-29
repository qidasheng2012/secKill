package controller

import (
	"log"
	"net/http"
	"seckill/admin/common"
	"seckill/admin/model"
	"seckill/admin/service"
	"seckill/admin/util"
	"strconv"
)

type SeckillController struct {
}

var seckillService = new(service.SeckillService)

func (s *SeckillController) Insert(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	numberStr := r.PostFormValue("number")
	startTimeStr := r.PostFormValue("startTime")
	endTimeStr := r.PostFormValue("endTime")

	if util.Empty(name) {
		common.ResultFail(w, "seckill name can not be empty")
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		common.ResultFail(w, "seckill number can not convert int")
		return
	}

	starterTime, e := util.ToTime(startTimeStr)
	if e != nil {
		common.ResultFail(w, "seckill startTime can not convert int")
		return
	}

	endTime, er := util.ToTime(endTimeStr)
	if er != nil {
		common.ResultFail(w, "seckill endTime can not convert int")
		return
	}

	seckill := model.Seckill{
		Name:      name,
		Number:    uint32(number),
		StartTime: starterTime,
		EndTime:   endTime,
	}

	id, err := seckillService.Insert(&seckill)

	if err != nil || id <= 0 {
		log.Printf("seckill Insert fail: %v", err)
		common.ResultFail(w, "seckill Insert fail")
		return
	}

	common.ResultOk(w, "seckill Insert success")
}

func (s *SeckillController) Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	seckillIdStr := query.Get("seckillId")
	id, err := strconv.Atoi(seckillIdStr)
	if err != nil {
		common.ResultFail(w, "seckill Delete fail, seckillId can not convert int")
		return
	}

	seckillId := int64(id)
	b := seckillService.Delete(seckillId)

	if !b {
		common.ResultFail(w, "seckill Delete fail")
		return
	}

	common.ResultOk(w, "seckill Delete ok")
}

func (s *SeckillController) Update(w http.ResponseWriter, r *http.Request) {
	seckillIdStr := r.PostFormValue("seckillId")
	name := r.PostFormValue("name")
	numberStr := r.PostFormValue("number")
	startTimeStr := r.PostFormValue("startTime")
	endTimeStr := r.PostFormValue("endTime")

	seckillId, er := strconv.Atoi(seckillIdStr)
	if er != nil {
		common.ResultFail(w, "seckillId can not convert int")
		return
	}

	if util.Empty(name) {
		common.ResultFail(w, "seckill name can not be empty")
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		common.ResultFail(w, "seckill number can not convert int")
		return
	}

	starterTime, e := util.ToTime(startTimeStr)
	if e != nil {
		common.ResultFail(w, "seckill startTime can not convert int")
		return
	}

	endTime, er := util.ToTime(endTimeStr)
	if er != nil {
		common.ResultFail(w, "seckill endTime can not convert int")
		return
	}

	seckill := model.Seckill{
		SeckillId: int64(seckillId),
		Name:      name,
		Number:    uint32(number),
		StartTime: starterTime,
		EndTime:   endTime,
	}

	err = seckillService.Update(&seckill)

	if err != nil {
		common.ResultFail(w, "seckill update fail")
		return
	}

	common.ResultOk(w, "seckill update success")
}

func (s *SeckillController) SelectById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	seckillIdStr := query.Get("seckillId")
	id, err := strconv.Atoi(seckillIdStr)
	if err != nil {
		common.ResultFail(w, "seckill SelectById fail, seckillId can not convert int")
		return
	}

	seckillId := int64(id)
	seckill, e := seckillService.SelectById(seckillId)

	if e != nil {
		common.ResultFail(w, "product SelectById error")
		return
	}

	common.ResultJsonOk(w, *seckill)
}

func (s *SeckillController) SelectAll(w http.ResponseWriter, r *http.Request) {
	seckills, err := seckillService.SelectAll()

	if err != nil {
		common.ResultFail(w, "seckill SelectAll error")
		return
	}

	common.ResultJsonOk(w, seckills)
}
