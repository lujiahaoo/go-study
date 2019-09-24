package main

import (
	// "fmt"
	// "runtime/debug"
	// "strings"
	"io"
	"log"
	"net/http"

	"github.com/lujiahaoo/go-study/routers"
	//"gopkg.in/gomail.v2"
	// "math/rand"
	// "flag"
	// "github.com/lujiahaoo/go-study/router"
	// "flag"
)

// func init()  {
// 	flag.StringVar(&name, "name", "everyone", "object")
// }

// var name *string;
// name=new(string)

// func test(i int) (j string) {
// 	defer func() {
// 		if p := recover(); p != nil {
// 			i = 1 - i
// 			fmt.Println(p)
// 			j = "ssss"
// 			DebugStack := ""
// 				for _, v := range strings.Split(string(debug.Stack()), "\n") {
// 					DebugStack += v + "<br>"
// 				}
// 			fmt.Println(DebugStack)

// 	}
// }()
// 	if i <= 0 {
// 		panic("please input int > 0")
// 	}
// 	return j
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func main() {
	// fmt.Println(test(-1))
	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", helloHandler)
	// mux.HandleFunc("/", echoHandler)

	// http.ListenAndServe("127.0.0.1:8080", mux)

	router := routers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

// func Count(ch chan int) {
//     fmt.Println("Counting")
//     ch <- 1
// }

// func main() {

//     chs := make([] chan int, 10)

//     for i:=0; i<10; i++ {
//         chs[i] = make(chan int)
//         go Count(chs[i])
//     }

//     for _,ch := range(chs) {
// 		<-ch
//     }
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main()  {
// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please input your name:")

// 	input, err := inputReader.ReadString('\n')
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		os.Exit(1)
// 	} else {
// 		name := input[0:len(input)-1]
// 		fmt.Printf("Hello, %s! what can i do for you ?\n", name)
// 	}

// 	for {
// 		input, err := inputReader.ReadString('\n')
// 		if err != nil {
// 			fmt.Printf("%s", err)
// 			continue
// 		}
// 		input = input[0:len(input)-1]

// 		input = strings.ToLower(input)

// 		switch input {
// 		case "":
// 			continue
// 		case "nothing", "bye":
// 			fmt.Println("Bye!")
// 			os.Exit(0)
// 		default:
// 			fmt.Printf("sorry,I did't catch you !+++ %s", input)
// 		}
// 	}
// }
