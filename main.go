package main

import (
	"fmt"
	"github.com/pocket52/ggame"
)

func main() {
	fmt.Println("Testing Pocket52")
	d := ggame.NewDeck()
	d.Print()
	fmt.Println(d.ToString())
	fmt.Println("No of Cards ", d.NoOfCards())
	fmt.Println(d.Take().String())
	fmt.Println("No of Cards ", d.NoOfCards())
	fmt.Println(d.Take().String())
	fmt.Println(d.Take().String())
	fmt.Println(d.Take().String())
	fmt.Println("No of Cards ", d.NoOfCards())

}
