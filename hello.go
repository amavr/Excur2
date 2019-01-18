package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	fmt.Printf("hello, world 2\n")
// }

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	// err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}
