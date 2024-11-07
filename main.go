package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/admin":
		fmt.Fprint(w, "I hate team18")
	default:
		fmt.Fprint(w, "404 File not found")
	}

}

func main() {
	http.HandleFunc("/", helloWorldPage)
	fmt.Println("Listening on port 80")
	err := http.ListenAndServe("", nil)
	if err != nil {
		fmt.Println("Error! message:", err)
	}
}
