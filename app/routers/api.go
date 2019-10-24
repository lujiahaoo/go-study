package routers

import (
	"net/http"

	"github.com/lujiahaoo/go-study/app/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//注册api前缀路由
func RegisteredApiRoute() Routes {
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

//注册normal前缀路由
func RegisteredNormalRoute() Routes {
	routes := Routes{
		Route{
			"Login",
			"POST",
			"/login",
			controllers.Login,
		},
		Route{},
		Route{
			"Register",
			"POST",
			"/register",
			controllers.Register,
		},
	}
	return routes
}
