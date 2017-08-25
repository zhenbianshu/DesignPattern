package main

import "builder"

func main() {
	chinese_style := builder.ChineseType{}
	chinese_builder := builder.NewBuilder(chinese_style)
	chinese_builder.Build() // A building with golden roof,red gate

	italian_style := builder.ItalianType{}
	italian_builder := builder.NewBuilder(italian_style)
	italian_builder.Build() // A building with white round roof,white gate
}
