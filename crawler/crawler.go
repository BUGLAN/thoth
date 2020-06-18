package crawler

import (
	"fmt"

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

func NewCssSelectorRule(rule *Rule) *CssCrawler {
	return &CssCrawler{Rule: rule, Collector: colly.NewCollector()}
}

type Callback func(s string) error

// Crawler 通过rule开始去爬数据
func (css *CssCrawler) Crawler(callbacks ...Callback) *CssCrawler {
	css.Collector.OnHTML(css.Rule.Selector, func(element *colly.HTMLElement) {
		fmt.Println("-----------------------------")
		s, err := element.DOM.Html()
		if err != nil {
			css.Error = err
		}
		fmt.Println(s)
		for _, v := range callbacks {
			err := v(s)
			if err != nil {
				css.Error = err
				return
			}
		}
	})

	css.Collector.OnResponse(func(response *colly.Response) {
		fmt.Println(response.StatusCode)
	})

	css.Collector.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	css.Error = css.Collector.Visit(css.Rule.Target)
	return css
}

// 假想api

// engine := crawler.New()
// engine.SetupRule(&Rule)
// engine.Start()
//
