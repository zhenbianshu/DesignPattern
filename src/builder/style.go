package builder

type style interface {
	roof() string
	gate() string
}

type ChineseType struct {
}

func (ChineseType) roof() string {
	return "golden roof"
}

func (ChineseType) gate() string {
	return "red gate"
}

type ItalianType struct {
}

func (ItalianType) roof() string {
	return "white round roof"
}

func (ItalianType) gate() string {
	return "white gate"
}
