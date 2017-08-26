package memento

type person struct {
	name   string
	age    int
	height int
	weight int
}

func (person *person) GetAge() int {
	return person.age
}

func (person *person) GrowUp(age, height, weight int) {
	person.age = age
	person.height = height
	person.weight = weight
}

func (person *person) Bak() memo {
	memo := memo{person.age, person.height, person.weight}

	return memo
}

func (person *person) Renewal(memo memo) {
	person.age = memo.age
	person.height = memo.height
	person.weight = memo.weight
}

func NewPerson(name string, age, height, weight int) *person {
	person := &person{name: name, age: age, height: height, weight: weight}

	return person
}
