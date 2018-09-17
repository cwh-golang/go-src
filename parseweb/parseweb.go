package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	SavePath  string
	MinWidth  int
	MinHeight int
	Overwrite bool
	MaxPage   int
	StartPage int
}

func NewConfig(savePath string, minWidth, minHeight, maxPage, startPage int, overwrite bool) *Config {
	return &Config{
		savePath,
		minWidth,
		minHeight,
		overwrite,
		maxPage,
		startPage,
	}
}

const (
	//https://pic.sogou.com/pics/recompic/detail.jsp?category=%E7%BE%8E%E5%A5%B3&tag=%E8%BF%B7%E4%BA%BA#0%2612638536
	PageUrl         string = "http://www.meizitu.com/a/sifang_5_%d.html"
	ImageListLinks  string = "http://www.meizitu.com/a/[0-9]+.html"
	ImageImageLinks string = "http://pic.meizitu.com/wp-content/uploads/[^\\.]+\\.(jpg|png|gif)"
)

type Webpage struct {
	Config *Config
}

func NewWebpage(config *Config) *Webpage {
	return &Webpage{Config: config}
}

func (wp *Webpage) ParsePage(url string) []string {
	offset := wp.Config.StartPage + wp.Config.MaxPage
	var urls []string
	for curPage := wp.Config.StartPage; curPage < offset; curPage ++ {
		urls = append(urls, fmt.Sprintf(url, curPage))
	}
	return urls
}

func (wp *Webpage) Get(url string) (body string) {
	resp, ok := http.Get(url)
	if nil != ok {
		return ""
	}
	defer resp.Body.Close()
	str, ok := ioutil.ReadAll(resp.Body)
	if ok != nil {
		return ""
	}
	return string(str)
}

func (wp *Webpage) ParseUrl(url, pattern string) (links []string) {
	fmt.Println("Parse url ==>", url)
	//fmt.Println("depart")
	body := wp.Get(url)
	if "" == body {
		fmt.Println("empty body")
		return []string{}
	}
	fmt.Println("body not empty ")
	reg := regexp.MustCompile(pattern)
	links = reg.FindAllString(body, -1)
	fmt.Println("get links :")
	fmt.Println(links)
	return links
}

func (wp *Webpage) GetSaveName(url string) string {
	paths := strings.Split(url, "/")
	len := len(paths)
	fileName := wp.Config.SavePath + paths[len-4] + paths[len-3] + paths[len-2] + paths[len-1]
	return fileName
}

func (wp *Webpage) Download(urls []string) {
	for _, url := range urls {
		fmt.Println("Start download image from url ==>", url)
		fileName := wp.GetSaveName(url)
		if wp.FileExist(fileName) && !wp.Config.Overwrite {
			fmt.Println("Image already exists, skip download ==>", url)
			continue
		}
		body := wp.Get(url)
		if "" == body {
			continue
		}
		if !wp.CheckSize(body, wp.GetExt(url)) {
			fmt.Println("Image size too small, skip download ==>", url)
			continue
		}
		if !wp.SaveImage(body, fileName) {
			fmt.Println("Save image failed ==>", url)
		}
	}
}

func (wp *Webpage) SaveImage(body, name string) bool {
	f, ok := os.Create(name)
	if ok != nil {
		fmt.Println("open file error")
		return false
	}
	defer f.Close()
	if _, err := f.WriteString(body); err == nil {
		return true
	}
	return false
}

func (wp *Webpage) GetExt(url string) string {
	if url == "" {
		return ""
	}
	temp := strings.Split(url, ".")
	return temp[len(temp)-1]
}

func (wp *Webpage) CheckSize(body, ext string) bool {
	if wp.Config.MinWidth <= 0 && wp.Config.MinHeight <= 0 {
		return true
	}
	var iImage image.Image
	var ok error = errors.New("unknown image type")
	switch ext {
	case "jpg":
		iImage, ok = jpeg.Decode(strings.NewReader(body))
	case "png":
		iImage, ok = png.Decode(strings.NewReader(body))
	case "gif":
		iImage, ok = gif.Decode(strings.NewReader(body))
	default:
		fmt.Println("unknown image format")
		return false
	}
	if ok == nil {
		rect := iImage.Bounds()
		if wp.Config.MinWidth <= rect.Max.X && wp.Config.MinHeight <= rect.Max.Y {
			return true
		}
	}
	return false
}

func (wp *Webpage) FileExist(name string) bool {
	if _, ok := os.Stat(name); ok == nil {
		return true
	}
	return false
}

func (wp *Webpage) RunTask() {
	urls := wp.ParsePage(PageUrl) //将要爬的网页url先记下来
	//fmt.Println(urls)
	//return
	sum := 0
	l := len(urls)
	c := make(chan int, l)

	for _, url := range urls {  //找到每个网页上的url
		go func(url string) {
			links := wp.ParseUrl(url, ImageListLinks)
			for _, v := range links {
				uris := wp.ParseUrl(v, ImageImageLinks)
				wp.Download(uris)
			}
			c <- 1
		}(url)
	}

forEnd:
	for {
		select {
		case <-c:
			sum ++
			if sum == l {
				break forEnd
			}
		}
	}

}

func main() {
	config := NewConfig(
		"F:\\girls\\",
		400,
		400,
		10,
		1,
		false,
	)

	webpage := NewWebpage(config)
	webpage.RunTask()

	fmt.Println("done!")
}
