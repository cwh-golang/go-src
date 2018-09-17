package realre

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Reretriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Reretriever) Get(url string) string {
	resp, err := http.Get(url)
	if err!=nil {
		panic(err)
	}

	bytes, e := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if e!=nil {
		panic(e)
	}
	return string(bytes)
}


