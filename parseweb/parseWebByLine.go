package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var filePath = "D:/download/parsePictures/"

func fileExist(fileName string) bool {
	if _, ok := os.Stat(fileName); ok == nil {
		return true
	}
	return false
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func SubStrTest() {
	s := Substr("123.jpg", -3, 4)
	fmt.Println(s)
}

func getImageList(url string, c chan int) {
	fmt.Println("get page link url==>", url)
	body := getUrl(url)
	if body == "" {
		fmt.Println("body empty:", url)
		return
	}
	fmt.Println("body not empty.")
	//fmt.Println("body:")
	//fmt.Println(body)
	//return
	reg := regexp.MustCompile("http://www.meizitu.com/a/[a-z]+.html")
	links := reg.FindAllString(body, -1)
	fmt.Println("all links:")
	//fmt.Println(links)
	fmt.Println("number of links :", len(links))
	//return
	getImageLink(links, c)
}

func getImageListTest() {
	//c := chan int
	//getImageList()
}

func getImageLink(links []string, c chan int) {
	for _, uri := range links {
		fmt.Println("Get images url, page link==>", uri)
		body := getUrl(uri)
		if "" == body {
			fmt.Println("empty body :", uri)
			return
		}
		fmt.Println(body)
		err := ioutil.WriteFile("body.txt", []byte(body), 0777)
		if err !=nil{
			fmt.Println("body.txt write f")
		}
		//reg := regexp.MustCompile("http://pic.meizitu.com/wp-content/uploads/[^\\.]+\\.(jpg|png|gif)")
		//reg := regexp.MustCompile("http ..(jpg|png|gif)")
		reg := regexp.MustCompile(`[a-zA-z]+://[^\s]*(jpg|png|gif|jpeg|JPG|PNG|JPEG)`)
		images := reg.FindAllString(body, -1)
		fmt.Printf("all images of %s:\n", uri)
		//fmt.Println(images)
		for _,v :=range images{
			fmt.Println(v)
		}
		fmt.Printf("number of iamges of %s is %d\n", uri, len(images))
		//continue
		downloadImage(images)
	}
	c <- 1
}

func getImageLinkTest() {
	pic := make([]string, 1)
	pic[0] = "http://gd7j.com/%E4%B9%89%E5%85%B4%E4%B9%A1%E9%A3%8E%E6%99%AF/"
	c:=make(chan int,1)
	getImageLink(pic, c)
}

func downloadImage(images []string) {
	for _, v := range images {
		fmt.Println("Download image, url==>", v)
		imageType := Substr(v, -2, 3)
		resp, ok := http.Get(v)
		if nil != ok {
			continue
		}
		defer resp.Body.Close()
		flag := false
		var iImage image.Image
		content, ok := ioutil.ReadAll(resp.Body)
		body := string(content)
		if imageType == "jpg" {
			iImage, ok = jpeg.Decode(strings.NewReader(body))
			flag = true
			fmt.Println("it is a jpg picture")
			if nil != ok {
				continue
			}
		} else if imageType == "png" {
			iImage, ok = png.Decode(strings.NewReader(body))
			flag = true
			fmt.Println("it is a png picture")
			if nil != ok {
				continue
			}
		} else if imageType == "jpeg" {
			iImage, ok = jpeg.Decode(strings.NewReader(body))
			flag = true
			fmt.Println("it is a jpeg picture")
			if nil != ok {
				continue
			}
		}

		if flag {
			rect := iImage.Bounds()
			if rect.Max.X < 200 || rect.Max.Y < 200 {
				//只下载大图，小图跳过
				fmt.Println("Skip download image, url ==>", v)
				continue
			}
		}

		// body:=getUrl(v)
		if nil != ok || "" == body {
			fmt.Println("content is empty")
			continue
		}
		paths := strings.Split(v, "/")
		len := len(paths)
		fileName := filePath + paths[len-4] + paths[len-3] + paths[len-2] + paths[len-1]
		fmt.Println("fileName of the picture is:", fileName)

		if fileExist(fileName) {
			fmt.Println("filename is exits")
			continue
		}

		f, ok := os.Create(fileName)
		if ok != nil {
			fmt.Println("open file error,the error is :", ok)
			return
		}
		defer f.Close()
		_, err := f.WriteString(body)
		if err != nil {
			fmt.Println("picture write fault")
		}
		fmt.Println("picture write succesfully")
	}
}

func downloadImageTest() {
	pic := make([]string, 1)
	pic[0] = "http://img3.xiazaizhijia.com/walls/20160929/mid_80815ca39dc3a63.jpg"
	//pic = append(pic, "http://seopic.699pic.com/photo/50045/5378.jpg_wh1200.jpg")
	for i, p := range pic {
		fmt.Println(i, p)
	}
	downloadImage(pic)
	fmt.Println()
}

func getUrl(url string) string {
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

func main() {
	//downloadImageTest()
	getImageLinkTest()
	os.Exit(0)


	fms := "http://www.meizitu.com/a/sifang_5_%d.html"
	ch := make(chan int, 1)

	url := fmt.Sprintf(fms, 1)
	fmt.Println("Parse url:", url)
	getImageList(url, ch)

	fmt.Println("done!")

}
