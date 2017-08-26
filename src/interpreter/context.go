package interpreter

type context struct {
	args map[string]expression
}

func (context *context) AddItem(key string, val expression) {
	context.args[key] = val
}

func NewContext() *context {
	args := make(map[string]expression) // map必须初始化后才能添加值
	context := &context{args}

	return context
}
