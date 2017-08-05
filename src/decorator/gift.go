package decorator

type Goods interface {
	GetDesc() string
}

type gift struct {
	desc string
}

func NewGift(name string) gift {
	gift := gift{name}
	return gift
}

func (gift gift) GetDesc() string {
	return gift.desc
}
