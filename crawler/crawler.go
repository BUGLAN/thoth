package crawler

import (
	"fmt"

	"github.com/axgle/mahonia"
	"github.com/gocolly/colly/v2"
)

type Rule struct {
	Target   string `json:"target"`
	Selector string
}

// CssCrawler ...
type CssCrawler struct {
	Rule      *Rule
	Error     error
	Tag       string
	Collector *colly.Collector
}

// Context crawler上下文
type Context struct {
	OriginUrl string
	Selector  string
	Html      string
	Text      string
	Element   *colly.HTMLElement
}

func NewCssSelectorRule(rule *Rule) *CssCrawler {
	return &CssCrawler{Rule: rule, Collector: colly.NewCollector()}
}

type Process func(ctx *Context) error

// Crawler 通过rule开始去爬数据
func (css *CssCrawler) Crawler(process Process) *CssCrawler {
	css.Collector.OnHTML(css.Rule.Selector, func(element *colly.HTMLElement) {
		s, err := element.DOM.Html()
		if err != nil {
			css.Error = err
		}

		css.Error = process(&Context{
			OriginUrl: css.Rule.Target,
			Selector:  css.Rule.Selector,
			Html:      s,
			Text:      element.DOM.Text(),
			Element:   element,
		})
	})

	// css.Collector.OnResponse(func(response *colly.Response) {
	//
	// })

	css.Collector.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	css.Error = css.Collector.Visit(css.Rule.Target)
	return css
}

// ConvertToString 编码转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
