package crawler

import (
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

func NewCssSelectorRule(rule *Rule) *CssCrawler {
	return &CssCrawler{Rule: rule, Collector: colly.NewCollector()}
}

type Callback func(html string, text string) error

// Crawler 通过rule开始去爬数据
func (css *CssCrawler) Crawler(callbacks ...Callback) *CssCrawler {
	css.Collector.OnHTML(css.Rule.Selector, func(element *colly.HTMLElement) {
		s, err := element.DOM.Html()
		if err != nil {
			css.Error = err
		}

		for _, v := range callbacks {
			err := v(s, element.Text)
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

func (*CssCrawler) ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
