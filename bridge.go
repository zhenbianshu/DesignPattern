package main

import "bridge"

func main() {
	pencil := bridge.Pencil{}
	ball_pen := bridge.BallPen{}

	red_paper := bridge.RedPaper{ball_pen}
	red_paper.Paint() // ballPen write on red paper!

	blue_paper := bridge.BluePaper{pencil}
	blue_paper.Paint() // pencil write on blue paper!
}
