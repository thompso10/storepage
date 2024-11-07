package main

import (
	"fmt"
	"net/http"
	"time"
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

func helloWorlHtml(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello World")
	w.Header().Set("Content-Type", "text/html")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello World</h1>")
	case "/admin":
		fmt.Fprint(w, "<h1>I hate team18</h1>")
	default:
		fmt.Fprint(w, "<h1>404 File not found</h1>")
	}
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Request Timeout failed :(")
}

func main() {
	http.HandleFunc("/", helloWorlHtml)
	http.HandleFunc("/timeout", timeout)

	fmt.Println("Listening on port 80")

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()

	// err := http.ListenAndServe("", nil)
	// if err != nil {
	// 	fmt.Println("Error! message:", err)
}
