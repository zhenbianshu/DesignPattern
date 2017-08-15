package main

import "adapter"

func main() {
	zhangsan := adapter.Person{"hello"}
	zhangsan.Express() // hello

	lisi := adapter.Mute{"I want to express!"}
	notebook := adapter.NoteBook{lisi}
	notebook.Express() // I want to express!
}
