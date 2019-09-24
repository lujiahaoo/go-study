package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	// w.Header().Set("Content-Type", "application/json: charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fmt.Fprintln(w, "the uid is:", uid)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fmt.Fprintln(w, "the user of ", uid, "is add")
}
