package mock

import (
	"fmt"
)

//Retriever
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("the Contents is %s",r.Contents)
}

func (r *Retriever)Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Read(p []byte) (n int, err error){
	fmt.Println(string(p))
	return len(p),nil
}

