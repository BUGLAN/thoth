package crawler

func ExampleCssSelectorRule_Crawler() {

	crawler := NewCssSelectorRule(&Rule{
		Target:   "http://www.shuquge.com/txt/107034/31841459.html",
		Selector: "#content",
	})
	crawler.Crawler()
	// OutPut:
}
