package bridge

type pen interface {
	write() string
}

type Pencil struct {
}

func (Pencil) write() string {
	return "pencil"
}

type BallPen struct {
}

func (BallPen) write() string {
	return "ballPen"
}
