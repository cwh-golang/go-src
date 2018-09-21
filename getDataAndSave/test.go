package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	_ "github.com/mysql"
)

func GetUrl(url string) string {
	resp, ok := http.Get(url)
	if nil != ok {
		fmt.Println(ok)
		return ""
	}
	defer resp.Body.Close()
	str, ok := ioutil.ReadAll(resp.Body)
	if ok != nil {
		return ""
	}
	return string(str)
}

func parseUrl(url string) []string {
	body := GetUrl(url)
	if body == "" {
		fmt.Printf("can not get info from %s", url)
		return nil
	}
	reg := regexp.MustCompile(`[a-zA-z]+://[^\s"]*`)
	return reg.FindAllString(body, -1)
}

const urlName = "http://www.baidu.com"

type urlDbStruct struct {
	Id  int
	Url string
}

func saveUrlToMysql(urls []string) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/mytest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 11
	for i, url := range urls {
		i += id
		sqlStr := fmt.Sprintf("insert into mytest.urltb(url) VALUES ('%s')", url)

		rs, err := db.Exec(sqlStr)
		if err != nil {
			log.Fatalln(err)
		}
		rowCount, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("inserted %d rows. \n", rowCount)
	}
}

func showURlFromMysql() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/mytest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM mytest.urltb")
	if err != nil {
		log.Fatalln(err)
	}
	num := 0
	for rows.Next() {
		u := urlDbStruct{
			Id:  0,
			Url: "",
		}
		err = rows.Scan(&u.Id,&u.Url)
		if err != nil {
			log.Fatalln(err)
		}
		num++
		fmt.Printf("query db to get id = %d ,the url = %s \n", u.Id,u.Url)
	}
	fmt.Println("total rows is :", num)
	rows.Close()
}

func main() {
	urls := parseUrl(urlName)
	for i,url := range urls {
		fmt.Println(i,"     ",url)
	}
	fmt.Println("total url is :",len(urls))
	saveUrlToMysql(urls)
	showURlFromMysql()
}
