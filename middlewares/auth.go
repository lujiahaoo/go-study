package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lujiahaoo/go-study/utils"
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
				fmt.Println(claminsToken.VerifyExpiresAt(time.Now().Unix(), true))
				//把 audience id 加到request中传递下去
				fmt.Println(claminsToken["aud"])
				next.ServeHTTP(w, r)
			}
		}
	})
}
