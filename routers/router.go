package routers

import (
	"github.com/gorilla/mux"
	"github.com/lujiahaoo/go-study/config"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//以下全部使用了全局中间件Cross
	config.RegisterGlobalMiddleWares(router)

	//基础路由，没用到中间件
	normal := router.PathPrefix("/").Subrouter()
	normalroutes := RegisteredNormalRoute()
	for _, route := range normalroutes {
		normal.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	//使用了路由中间件Auth，并且匹配了/api才可以访问子路由
	api := router.PathPrefix("/api").Subrouter()
	config.RegisterGroupMiddleWares(api)
	apiroutes := RegisteredApiRoute() //routes.go文件的注册路由函数,添加路由在里面写
	for _, route := range apiroutes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
