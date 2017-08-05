package observer

import "fmt"

// 记者接口，有报道方法
type IReporter interface {
	report(string)
}

// 新闻记者基类， 有名字属性和并实现了报道方法
type newsReporter struct {
	name string
}

func NewNewsReporter(name string) newsReporter {
	news_reporter := newsReporter{name}

	return news_reporter
}

func (reporter newsReporter) report(event string) {
	fmt.Println(reporter.name + " reported an event:" + event)
}
