package main

import (
	"os"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is version 3")
}

func main() {
	http.HandleFunc("/", helloHandler)

	var port string = "8080"
	argLength := len(os.Args[1:])
	fmt.Println("Found ", argLength, " args")
	if argLength != 0 {
		port = os.Args[1]
	}
	fmt.Println("Port is " + port)

	fmt.Println("Started, serving at "+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
