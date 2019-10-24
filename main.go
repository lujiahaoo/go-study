package main

import (
	// "fmt"
	// "runtime/debug"
	// "strings"
	// "math/rand"
	// "flag"
	"log"
	"net/http"

	"github.com/lujiahaoo/go-study/app/routers"
	//"gopkg.in/gomail.v2"
	// "github.com/lujiahaoo/go-study/router"
)

func main() {
	// fmt.Println(test(-1))
	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", helloHandler)
	// mux.HandleFunc("/", echoHandler)

	// http.ListenAndServe("127.0.0.1:8080", mux)

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
