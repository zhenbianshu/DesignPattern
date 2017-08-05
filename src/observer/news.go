package observer

type News struct {
	reporters []IReporter
}

func (news News) Happen(event string) {
	for _, reporter := range news.reporters {
		reporter.report(event)
	}
}

func (news *News) Focus(reporter IReporter) {
	news.reporters = append(news.reporters, reporter)
}

func (news *News) Ignore(reporter IReporter) {
	for index, news_reporter := range news.reporters {
		if news_reporter == reporter {
			news.reporters = append(news.reporters[:index], news.reporters[index+1:]...)
		}
	}
}
