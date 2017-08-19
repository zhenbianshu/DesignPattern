package proxy

type lazyPostman struct {
	letter letterContent
}

func NewLazyPostman() lazyPostman {
	letter := tempContent{}
	return lazyPostman{letter}
}

func (postman lazyPostman) RequestLetter() {
	postman.letter.get()
}

func (postman *lazyPostman) DemandLetter() {
	postman.letter = realContent{}
	postman.letter.get()
}
