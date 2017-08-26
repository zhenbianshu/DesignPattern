package mediator

type unitedNations struct {
	europe *country
	asia   *country
}

func (mediator unitedNations) getMsg(country *country) string {
	if mediator.europe == country {
		return mediator.asia.msg
	} else {
		return mediator.europe.msg
	}
}

func NewUnitedNations(european_country, asian_countryB *country) *unitedNations {
	unitedNations := &unitedNations{european_country, asian_countryB}

	return unitedNations
}
