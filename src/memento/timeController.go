package memento

type memo struct {
	age    int
	height int
	weight int
}

type timeController struct {
	memos map[int]memo
}

func NewTimeController() *timeController {
	memos := make(map[int]memo)
	controller := &timeController{memos}

	return controller
}

func (controller *timeController) Save(age int, memo memo) {
	controller.memos[age] = memo
}

func (controller *timeController) GetMemo(age int) memo {
	return controller.memos[age]
}
