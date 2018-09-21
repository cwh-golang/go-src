package main

import (
	"fmt"
	"regexp"
)

func parseUrl(url string) []string {
	body := GetUrl(url)
	if body == "" {
		fmt.Printf("can not get info from %s", url)
		return nil
	}
	reg := regexp.MustCompile(`[a-zA-z]+://[^\s"]*`)
	return reg.FindAllString(body, -1)
}

const urlName  = "www.runoob.com/mysql/mysql-create-database.html"

func main() {
	urls := parseUrl(urlName)
	fmt.Println(urls)

	//saveUrl()
	//showURl()
}
