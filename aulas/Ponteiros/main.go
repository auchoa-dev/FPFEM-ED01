package main

import (
	"fmt"
)

func main() {
	var dolly int
	dolly = 10

	var dollynha int
	dollynha = 30

	var p *int
	p = &dolly

	*p = dollynha
	fmt.Println("1 - ", dolly)
	fmt.Println("2 - ", dollynha)
	fmt.Println("3 - ", *p)

}
