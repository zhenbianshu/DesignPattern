package main

import "observer"

func main() {
	news := observer.News{}

	zhangsan := observer.NewNewsReporter("zhangsan")
	news.Focus(zhangsan)

	lisi := observer.NewNewsReporter("lisi")
	news.Focus(lisi)

	news.Happen("someone was killed!") // zhangsan reported an event:someone was killed!  lisi reported an event:someone was killed!

	news.Ignore(lisi)
	news.Happen("a building built in the city!") // zhangsan reported an event:a building built in the city!
}
