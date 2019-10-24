package routers

import (
	"github.com/gorilla/mux"
	"github.com/lujiahaoo/go-study/bootstrap"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//以下全部使用了全局中间件Cross
	bootstrap.RegisterGlobalMiddleWares(router)

	//基础路由，除了全局中间件没用到其他中间件
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
	//中间件需要注册后才可使用
	bootstrap.RegisterGroupMiddleWares(api)
	//routes.go文件的注册路由函数,要添加路由都在里面写
	apiroutes := RegisteredApiRoute()
	for _, route := range apiroutes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
