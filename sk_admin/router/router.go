package router

import (
	"SecKill/sk_admin/common"
	"SecKill/sk_admin/controller"
)

func RegiterRouter(handler *common.RouterHandler) {
	// 商品相关路由
	productController := new(controller.ProductController)

	handler.Router("/product/insert", productController.Insert)
	handler.Router("/product/delete", productController.Delete)
	handler.Router("/product/update", productController.Update)
	handler.Router("/product/selectById", productController.SelectById)
	handler.Router("/product/selectAll", productController.SelectAll)

	// 秒杀商品信息路由

}
