package iterator

type soldiers struct {
	queue  map[string]string
	names  []string // map is not ordered, use a slice to store the order
	cursor int
}

func NewSoldiers(queue map[string]string) *soldiers {
	var names []string
	for index := range queue {
		names = append(names, index)
	}
	soldiers := &soldiers{queue, names, -1}

	return soldiers
}

func (soldiers soldiers) Count() int {
	return len(soldiers.queue)
}

func (soldiers *soldiers) Next() string {

	if soldiers.cursor >= len(soldiers.names) {
		return "out of range!"
	} else {
		soldiers.cursor++
		return soldiers.queue[soldiers.names[soldiers.cursor]]
	}
}

func (soldiers *soldiers) Remove() {
	for index, name := range soldiers.names {
		if index == soldiers.cursor {
			soldiers.names = append(soldiers.names[:index], soldiers.names[index+1:]...)
			delete(soldiers.queue, name)
		}
	}
}
