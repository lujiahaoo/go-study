package routers

import (
	"net/http"

	"github.com/lujiahaoo/go-study/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//注册路由
func RegisteredRoute() Routes {
	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			controllers.Index,
		},
		Route{
			"AllUser",
			"GET",
			"/user",
			controllers.AllUser,
		},
		Route{
			"ShowUser",
			"GET",
			"/user/{uid:[0-9]+}",
			controllers.ShowUser,
		},
		Route{
			"AddUser",
			"POST",
			"/user/{uid:[0-9]+}",
			controllers.AddUser,
		},
	}
	return routes
}
