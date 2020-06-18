package crawler

func ExampleCssSelectorRule_Crawler() {

	crawler := NewCssSelectorRule(&Rule{
		Target:   "https://gist.github.com/BUGLAN/0187c5cda0e99dcd2824a0e3b1599ee4",
		Selector: "file-box",
	})
	crawler.Crawler()
	// OutPut:
}
