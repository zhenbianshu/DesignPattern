package bridge

import "fmt"

type paper interface {
	Paint()
}

type RedPaper struct {
	Pen pen
}

type BluePaper struct {
	Pen pen
}

func (paper RedPaper) Paint() {
	fmt.Println(paper.Pen.write() + " write on red paper!")
}

func (paper BluePaper) Paint() {
	fmt.Println(paper.Pen.write() + " write on blue paper!")
}
