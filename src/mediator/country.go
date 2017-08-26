package mediator

type country struct {
	name string
	msg  string
}

func (country *country) Receive(nations *unitedNations) string {
	return nations.getMsg(country)
}

func (country *country) Announce(msg string) {
	country.msg = msg
}

func NewCountry(name string) *country {
	country := &country{name: name}

	return country
}
