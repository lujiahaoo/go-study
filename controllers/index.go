package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lujiahaoo/go-study/untils"
)

/*假数据begin*/
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

/*假数据end*/

func Index(w http.ResponseWriter, R *http.Request) {
	fmt.Fprintln(w, "welcome to go-study")
}

func AllUser(w http.ResponseWriter, R *http.Request) {
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
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {3, 2, 1}}

	resp := untils.Resp{
		"200",
		"ok",
		arr,
	}
	resp.MarshalJson(w)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fmt.Fprintln(w, "the user of ", uid, "is add")
}
