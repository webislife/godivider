package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	var a string = req.URL.Query().Get("a")
	var b string = req.URL.Query().Get("b")

	if(len(a) == 0 || len(b) == 0) {
		return
	}
	af, err :=  strconv.ParseFloat(a, 64)
	bf, err :=  strconv.ParseFloat(b, 64)
	if(err != nil) {
		log.Fatal(err)
	}
	fmt.Fprintf(writer, strconv.FormatFloat(af/bf, 'g', 100, 64))
}

func main () {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9091", nil)
}