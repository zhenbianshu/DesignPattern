package state

type fruit struct {
	stage Stage
}

func NewFruit() fruit {
	seedling := seedling{}
	fruit := fruit{seedling}

	return fruit
}

func (fruit *fruit) Water() {
	fruit.stage.water()
	if _, ok := interface{}(fruit.stage).(seedling); ok {
		bloom := bloom{}
		fruit.stage = bloom
	} else {
		maturity := maturity{}
		fruit.stage = maturity
	}
}

func (fruit *fruit) Harvest() {
	fruit.stage.harvest()
	if _, ok := interface{}(fruit.stage).(maturity); ok {
		seedling := seedling{}
		fruit.stage = seedling
	}
}
