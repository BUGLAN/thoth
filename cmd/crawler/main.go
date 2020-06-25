package main

import (
	"errors"
	"flag"

	"github.com/BUGLAN/thoth/crawler"
	"github.com/BUGLAN/thoth/wire"
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
	process, err := wire.InitKanShuProcess()
	if err != nil {
		panic(err)
	}
	if err = cssCrawler.Crawler(process.GithubReadMeProcess).Error; err != nil {
		panic(err)
	}
}
