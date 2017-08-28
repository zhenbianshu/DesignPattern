package prototype

type honorGuard struct {
	age    int
	height int
	weight int
	gender string
	name   string
}

func (guard *honorGuard) SetName(name string) {
	guard.name = name
}

func (guard *honorGuard) Clone() *honorGuard {
	new_guard := *guard

	return &new_guard
}

func NewhonorGuard(age, height, weight int, gender string) *honorGuard {
	guard := &honorGuard{age: age, height: height, weight: weight, gender: gender}

	return guard
}
