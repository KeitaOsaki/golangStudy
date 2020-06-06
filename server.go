package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("hello.html"))

	if err := t.Execute(os.Stdout, nil); err != nil { // テンプレート出力
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
