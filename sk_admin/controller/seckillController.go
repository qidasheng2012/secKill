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
		common.ResultFail(w, "name can't be empty")
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		common.ResultFail(w, "number can't convert int")
		return
	}

	starterTime, e := util.ToTime(startTimeStr)
	if e != nil {
		common.ResultFail(w, "startTime can't convert time")
		return
	}

	endTime, er := util.ToTime(endTimeStr)
	if er != nil {
		common.ResultFail(w, "endTime can't convert time")
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
		common.ResultFail(w, "Insert fail")
		return
	}

	common.ResultOk(w, "Insert success")
}

func (s *SeckillController) Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	seckillIdStr := query.Get("seckillId")
	id, err := strconv.Atoi(seckillIdStr)
	if err != nil {
		common.ResultFail(w, "seckillId can't convert int")
		return
	}

	err = seckillService.Delete(int64(id))

	if err != nil {
		log.Printf("seckillId [%s]  Delete fail: %v ", seckillIdStr, err)
		common.ResultFail(w, "Delete fail")
		return
	}

	common.ResultOk(w, "Delete success")
}

func (s *SeckillController) Update(w http.ResponseWriter, r *http.Request) {
	seckillIdStr := r.PostFormValue("seckillId")
	name := r.PostFormValue("name")
	numberStr := r.PostFormValue("number")
	startTimeStr := r.PostFormValue("startTime")
	endTimeStr := r.PostFormValue("endTime")

	seckillId, er := strconv.Atoi(seckillIdStr)
	if er != nil {
		common.ResultFail(w, "seckillId can't convert int")
		return
	}

	if util.Empty(name) {
		common.ResultFail(w, "name can't be empty")
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		common.ResultFail(w, "number can't convert int")
		return
	}

	starterTime, e := util.ToTime(startTimeStr)
	if e != nil {
		common.ResultFail(w, "startTime can't convert int")
		return
	}

	endTime, er := util.ToTime(endTimeStr)
	if er != nil {
		common.ResultFail(w, "endTime can't convert int")
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
		common.ResultFail(w, "update fail")
		return
	}

	common.ResultOk(w, "update success")
}

func (s *SeckillController) SelectById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	seckillIdStr := query.Get("seckillId")
	id, err := strconv.Atoi(seckillIdStr)
	if err != nil {
		common.ResultFail(w, "seckillId can't convert int")
		return
	}

	seckillId := int64(id)
	seckill, e := seckillService.SelectById(seckillId)

	if e != nil {
		common.ResultFail(w, "SelectById error")
		return
	}

	common.ResultJsonOk(w, *seckill)
}

func (s *SeckillController) SelectEtcdById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	seckillIdStr := query.Get("seckillId")

	seckill, e := seckillService.SelectEtcdById(seckillIdStr)

	if e != nil {
		common.ResultFail(w, "selectEtcdById error")
		return
	}

	if seckill == nil {
		common.ResultFail(w, "selectEtcdById value is empty")
		return
	}

	common.ResultJsonOk(w, *seckill)
}

func (s *SeckillController) SelectAll(w http.ResponseWriter, r *http.Request) {
	seckills, err := seckillService.SelectAll()

	if err != nil {
		common.ResultFail(w, "SelectAll error")
		return
	}

	common.ResultJsonOk(w, seckills)
}
