package controller

import (
	"net/http"
	"strconv"

	"seckill/admin/common"
	"seckill/admin/model"
	"seckill/admin/service"
	"seckill/admin/util"
)

type ProductController struct {
}

var productService = new(service.ProductService)

func (p *ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	productName := r.PostFormValue("productName")
	totalStr := r.PostFormValue("total")
	statusStr := r.PostFormValue("status")

	if util.Empty(productName) {
		common.ResultFail(w, "productName can not be empty")
		return
	}

	total, err := strconv.Atoi(totalStr)
	if err != nil {
		common.ResultFail(w, "total can not convert int")
		return
	}

	status, e := strconv.Atoi(statusStr)
	if e != nil {
		common.ResultFail(w, "status can not convert int")
		return
	}

	product := model.Product{
		ProductName: productName,
		Total:       uint64(total),
		Status:      uint(status),
	}

	id, err := productService.Insert(&product)

	if err != nil || id <= 0 {
		common.ResultFail(w, "product Insert fail")
		return
	}

	common.ResultOk(w, "product Insert success")
}

func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	productIdStr := query.Get("productId")
	id, err := strconv.Atoi(productIdStr)
	if err != nil {
		common.ResultFail(w, "product Delete fail, productId can not convert int")
		return
	}

	productId := int64(id)
	b := productService.Delete(productId)

	if !b {
		common.ResultFail(w, "product Delete fail")
		return
	}

	common.ResultOk(w, "product Delete ok")
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	productIdStr := r.PostFormValue("productId")
	productName := r.PostFormValue("productName")
	totalStr := r.PostFormValue("total")
	statusStr := r.PostFormValue("status")

	productId, er := strconv.Atoi(productIdStr)
	if er != nil {
		common.ResultFail(w, "productId can not convert int")
		return
	}

	if util.Empty(productName) {
		common.ResultFail(w, "productName can not be empty")
		return
	}

	total, err := strconv.Atoi(totalStr)
	if err != nil {
		common.ResultFail(w, "total can not convert int")
		return
	}

	status, e := strconv.Atoi(statusStr)
	if e != nil {
		common.ResultFail(w, "status can not convert int")
		return
	}

	product := model.Product{
		ProductId:   uint64(productId),
		ProductName: productName,
		Total:       uint64(total),
		Status:      uint(status),
	}

	err = productService.Update(&product)

	if err != nil {
		common.ResultFail(w, "product update fail")
		return
	}

	common.ResultOk(w, "product update success")
}

func (p *ProductController) SelectById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	productIdStr := query.Get("productId")
	id, err := strconv.Atoi(productIdStr)
	if err != nil {
		common.ResultFail(w, "product SelectById fail, productId can not convert int")
		return
	}

	productId := int64(id)
	product, e := productService.SelectById(productId)

	if e != nil {
		common.ResultFail(w, "product SelectById error")
		return
	}

	common.ResultJsonOk(w, *product)
}

func (p *ProductController) SelectAll(w http.ResponseWriter, r *http.Request) {
	products, err := productService.SelectAll()

	if err != nil {
		common.ResultFail(w, "product SelectAll error")
		return
	}

	common.ResultJsonOk(w, products)
}
