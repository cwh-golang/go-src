package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world,你好世界`))
	})
	http.ListenAndServe(":3000", nil) // <-今天讲的就是这个ListenAndServe是如何工作的
}



//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "hello world")
//}
//
//func main() {
//	http.HandleFunc("/", IndexHandler)
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}
