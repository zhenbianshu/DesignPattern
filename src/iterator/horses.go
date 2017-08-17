package iterator

type horses struct {
	queue  []string
	cursor int
}

func NewHorses(queue []string) *horses {
	horses := &horses{queue, -1}

	return horses
}

func (horses horses) Count() int {
	return len(horses.queue)
}

func (horses *horses) Next() string {
	if horses.cursor >= len(horses.queue) {
		return "out of range!"
	} else {
		horses.cursor++
		return horses.queue[horses.cursor]
	}
}

func (horses *horses) Remove() {
	for index := range horses.queue {
		if index == horses.cursor {
			horses.queue = append(horses.queue[:index], horses.queue[index+1:]...)
		}
	}
}
