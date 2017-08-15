package singleton

var object *instance

type instance struct {
	name string
}

func (instance instance) GetName() string {
	return instance.name
}

func (instance *instance) SetName(name string) {
	instance.name = name
}

// todo 后续考虑并行问题
func GetInstance() *instance {
	if object == nil {
		object = &instance{"test1"}
	}

	return object
}
