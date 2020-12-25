package dao

import (
	"database/sql"
	"log"

	"SecKill/sk_admin/common"
	"SecKill/sk_admin/model"
)

type IProductDao interface {
	Insert(*model.Product) (int64, error)
	Delete(int64) bool
	Update(*model.Product) error
	SelectById(int64) (*model.Product, error)
	SelectAll() ([]*model.Product, error)
}

type ProductDao struct {
}

func (p *ProductDao) Insert(product *model.Product) (int64, error) {
	result, err := common.DB.Exec("INSERT INTO product(`product_name`,`total`,`status`) value(?,?,?)",
		product.ProductName, product.Total, product.Status)

	if err != nil {
		return 0, err
	}

	id, e := result.LastInsertId()
	if e != nil {
		return 0, e
	}
	return id, nil
}

func (p *ProductDao) Delete(productId int64) bool {
	_, err := common.DB.Exec("DELETE FROM product where product_id = ?", productId)
	if err != nil {
		return false
	}

	return true
}

func (p *ProductDao) Update(product *model.Product) error {
	_, err := common.DB.Exec("UPDATE product SET product_name = ? , total = ? , status = ? where product_id = ?",
		product.ProductName, product.Total, product.Status, product.ProductId)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductDao) SelectById(productId int64) (*model.Product, error) {
	rows, err := common.DB.Query("SELECT product_id,product_name,total,status from product where product_id = ?", productId)
	if err != nil {
		return nil, err
	}

	products, err := processRows(rows)
	if err != nil {
		return nil, err
	}

	return products[0], nil
}

func (p *ProductDao) SelectAll() ([]*model.Product, error) {
	rows, err := common.DB.Query("SELECT product_id,product_name,total,status from product")
	if err != nil {
		return nil, err
	}

	return processRows(rows)
}

func processRows(rows *sql.Rows) ([]*model.Product, error) {
	var products []*model.Product

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ProductId, &product.ProductName, &product.Total, &product.Status)
		if err != nil {
			log.Printf("product  SelectById error: %v", err)
			continue
		}
		products = append(products, &product)
	}
	rows.Close()

	return products, nil
}
