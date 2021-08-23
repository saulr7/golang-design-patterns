package main

import (
	"designpatterns/creational/singleton"
	"fmt"
)

func main() {

	counter := singleton.GetInstance()
	counter2 := singleton.GetInstance()

	fmt.Println(counter.AddOne())
	fmt.Println(counter.AddOne())
	fmt.Println(counter2.AddOne())

}
