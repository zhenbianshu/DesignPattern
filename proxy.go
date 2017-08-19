package main

import "proxy"

func main() {
	postman := proxy.NewLazyPostman()
	postman.RequestLetter() // I am try my best to get the reply letter.
	postman.DemandLetter()  // The content is:I love you, too.
}
