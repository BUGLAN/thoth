package crawler

import (
	"fmt"
)

// ExampleCssSelectorRule_Crawler 爬取小说网站
func ExampleCssSelectorRule_Crawler() {
	crawler := NewCssSelectorRule(&Rule{
		Target:   "https://gist.github.com/BUGLAN/0187c5cda0e99dcd2824a0e3b1599ee4",
		Selector: "file-box",
	})
	crawler.Crawler(func(ctx *Context) error {
		fmt.Println(ctx.Html)
		return nil
	})
	// OutPut:
}

func ExampleCssSelectorRule_Crawler1() {
	crawler := NewCssSelectorRule(&Rule{
		Target:   "https://m.38kanshu.cc/191854/mulu.html",
		Selector: ".chapter li",
	})
	crawler.Crawler(func(ctx *Context) error {
		s := ConvertToString(ctx.Html, "gbk", "utf-8")
		fmt.Println(s)
		return nil
	})
	// OutPut:
}
