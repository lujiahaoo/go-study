package routers

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/v1").Subrouter() //添加版本前缀,匹配了/v1才可以访问子路由
	routes := RegisteredRoute()                                               //routes.go文件的注册路由函数,添加路由在里面写
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
