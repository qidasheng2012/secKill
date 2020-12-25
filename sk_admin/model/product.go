package model

type Product struct {
	ProductId   uint64 `json:"productId"`
	ProductName string `json:"productName"`
	Total       uint64 `json:"total"`
	Status      uint   `json:"status"`
}
