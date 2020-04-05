package main

import (
	// "fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/index", contentFunc)
	// http.HandleFunc("/", contentFunc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func contentFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "2oops")
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}