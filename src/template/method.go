package template

type Sample interface {
	getSender(name string) string
	processContent(content string) string
	getReceiver(name string) string
}

type Sms struct {
}

func (sender Sms) getSender(name string) string {
	return name + "'s phone_number :110"
}

func (sender Sms) processContent(content string) string {
	return "sms:" + content
}

func (sender Sms) getReceiver(name string) string {
	return name + "'s phone_number: 12345"
}

type Mail struct {
}

func (sender Mail) getSender(name string) string {
	return name + "@126.com"
}

func (sender Mail) processContent(content string) string {
	return "mail:" + content
}

func (sender Mail) getReceiver(name string) string {
	return name + "@163.com"
}
