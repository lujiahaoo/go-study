package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/lujiahaoo/go-study/app/middlewares"
)

//注册全局中间件
func RegisterGlobalMiddleWares(router *mux.Router) {
	router.Use(
		middlewares.CrossOrigin,
		middlewares.Exception,
	)
}

//注册路由组中间件
func RegisterGroupMiddleWares(router *mux.Router) {
	router.Use(
		middlewares.AuthCheck,
	)
}
