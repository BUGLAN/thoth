package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"

	"github.com/BUGLAN/thoth/crawler"
)

var (
	url      string
	selector string
	err      error
)

func init() {
	flag.StringVar(&url, "u", "", "url链接地址")
	flag.StringVar(&selector, "s", "", "选择器")
}

func main() {
	flag.Parse()
	if url == "" || selector == "" {
		panic(errors.New("url 或 selector 不能为空"))
	}

	cssCrawler := crawler.NewCssSelectorRule(&crawler.Rule{
		Target:   url,
		Selector: selector,
	})

	if err = cssCrawler.Crawler(func(html string, text string) error {
		html = cssCrawler.ConvertToString(html, "gbk", "utf-8")
		compile := regexp.MustCompile(`<a href="(.*?)".*>(.*?)<span>`)

		stringList := compile.FindStringSubmatch(html)
		fmt.Println(stringList)
		if len(stringList) == 3 {
			fmt.Println(stringList[1], stringList[2])
		}

		return nil
	}).Error; err != nil {
		fmt.Println(err)
	}
}
