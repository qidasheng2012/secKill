package router

import (
	"seckill/admin/common"
	"seckill/admin/controller"
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
	seckillController := new(controller.SeckillController)
	handler.Router("/seckill/insert", seckillController.Insert)
	handler.Router("/seckill/delete", seckillController.Delete)
	handler.Router("/seckill/update", seckillController.Update)
	handler.Router("/seckill/selectById", seckillController.SelectById)
	handler.Router("/seckill/selectAll", seckillController.SelectAll)

}
