package main

import (
	"fmt"
	"net/http"
)

type Handler struct{}

var NewUrl string = "http://google.com"

func test(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, NewUrl, http.StatusFound)
	fmt.Fprintf(w, " ")

	return
}

func result(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<body>Duh</body>")
	fmt.Fprintf(w, "<br>localhost:8081/test/")
	return
}

func main() {
	http.HandleFunc("/test/", test)
	http.HandleFunc("/result/", result)
	http.ListenAndServe("0.0.0.0:9000", nil)
}
