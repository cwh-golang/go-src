package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter,
		request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		fmt.Println(path)
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		//fmt.Println(string(bytes))
		if err != nil {
			panic(err)
		}

		writer.Write(bytes)
	})

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	//err := http.ListenAndServe(":8888", nil)
	//err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		panic(err)
	}
}
