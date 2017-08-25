package flyweight

import (
	"fmt"
	"strconv"
)

type forest struct {
	trees []map[string]int
}

func (forest forest) Fell() {
	for _, attr := range forest.trees {
		fmt.Println("a " + strconv.Itoa(attr["height"]) + " m tree cutted!")
	}
}

func NewForest(trees []map[string]int) forest {
	forest := forest{trees}

	return forest
}
