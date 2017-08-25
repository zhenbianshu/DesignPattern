package builder

import "fmt"

type builder struct {
	style
}

func (builder builder) Build() {
	fmt.Println("A building with " + builder.roof() + "," + builder.gate())
}

func NewBuilder(style style) builder {
	building := builder{style: style}

	return building
}
