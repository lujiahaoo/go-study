package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/lujiahaoo/go-study/app/utils"
)

//实现身份验证的中间件
func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			payload := utils.Response{
				Code: http.StatusUnauthorized,
				Msg:  "not authorized",
			}
			utils.ResponseWithJson(w, payload, http.StatusUnauthorized)
		} else {

			//将token字符串就是类似nuisd2137.d21qy789712,.231y9sdjkad的这坨东西转换成token结构体
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				//验证Header alg加密算法是否合规
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					utils.ResponseWithJson(
						w,
						utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"},
						http.StatusUnauthorized)
					return nil, fmt.Errorf("not authorized")
				}
				return []byte(utils.ENV("appkey").(string)), nil
			})
			if !token.Valid {
				utils.ResponseWithJson(
					w,
					utils.Response{Code: http.StatusUnauthorized, Msg: "not authorized"},
					http.StatusUnauthorized)
			} else {
				claminsToken := token.Claims.(jwt.MapClaims)
				//判断是否过期(设置了iat 和 exp就不用这一步了，在.Valid中会自动验证)
				//fmt.Println(claminsToken.VerifyExpiresAt(time.Now().Unix(), true))

				//把 audience id 加到Header中传递下去,后面直接拿request Header中就有用户id
				r.Header.Set("sub", claminsToken["sub"].(string))
				next.ServeHTTP(w, r)
			}
		}
	})
}
