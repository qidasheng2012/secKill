package service

import (
	"github.com/seckill/admin/dao"
	"github.com/seckill/admin/model"
)

var productDao = new(dao.ProductDao)

type ProductService struct {
}

func (p *ProductService) Insert(product *model.Product) (int64, error) {
	return productDao.Insert(product)
}

func (p *ProductService) Delete(productId int64) bool {
	return productDao.Delete(productId)
}

func (p *ProductService) Update(product *model.Product) error {
	return productDao.Update(product)
}

func (p *ProductService) SelectById(productId int64) (*model.Product, error) {
	return productDao.SelectById(productId)
}

func (p ProductService) SelectAll() ([]*model.Product, error) {
	return productDao.SelectAll()
}
