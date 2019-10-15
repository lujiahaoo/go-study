package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lujiahaoo/go-study/utils"
)

/*假数据begin*/
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

/*假数据end*/

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to go-study")
}

func AllUser(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{
			Name:      "W P",
			Completed: true,
		},
		Todo{
			Name:      "S S",
			Completed: false,
		},
	}
	//加了下面的就变成下载文件了，有时间去查查
	// w.Header().Set("Content-Type", "application/json: charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	// var arr [2][3]int = [2][3]int{{1, 2, 3}, {3, 2, 1}}

	resp := utils.Response{
		Code: 200,
		Msg:  "ok",
	}
	utils.ResponseWithJson(w, resp)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fmt.Fprintln(w, "the user of ", uid, "is add")
}

// func GenerateToken(w http.ResponseWriter, r *http.Request) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	clamis := make(jwt.MapClaims)
// 	clamis["aud"] = "lu"
// 	exp := utils.ENV("exp").(int64) * 3600 * 24
// 	clamis["exp"] = time.Now().Unix() + exp
// 	clamis["iat"] = time.Now().Unix()
// 	token.Claims = clamis

// 	tokenString, _ := token.SignedString([]byte(utils.ENV("appkey").(string)))
// 	fmt.Println("my token is: ", tokenString)

// 	tokenParse, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//不要忘了验证Header的加密算法是否是你想要的
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(utils.ENV("appekey").(string)), nil
// 	})

// 	findToken := tokenParse.Claims.(jwt.MapClaims)
// 	resp := utils.Response{
// 		Code: 200,
// 		Msg:  "ok",
// 		Data: findToken["aud"],
// 	}
// 	utils.ResponseWithJson(w, resp)
// }
