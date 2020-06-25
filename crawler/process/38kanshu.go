/*
爬虫处理器, 将爬虫获取到的内容, 进行进一步解析, 去重, 处理, 存储,
*/
package process

import (
	"fmt"
	"regexp"

	"github.com/BUGLAN/thoth/crawler"
	"github.com/BUGLAN/thoth/dao"
)

type KanShu struct {
	book    dao.BookDao
	article dao.ArticleDao
}

func NewKanShu(book dao.BookDao, article dao.ArticleDao) *KanShu {
	return &KanShu{
		book:    book,
		article: article,
	}
}

// KanShuProcess 处理器
// https://m.38kanshu.cc/191854/mulu.html
func (process *KanShu) KanShuProcess(html string, text string) error {
	html = crawler.ConvertToString(html, "gbk", "utf-8")
	compile := regexp.MustCompile(`<a href="(.*?)".*>(.*?)<span>`)

	stringList := compile.FindStringSubmatch(html)
	if len(stringList) == 3 {
		fmt.Println(stringList[1], stringList[2])
	}
	// save
	// process.book.Create()

	return nil
}
