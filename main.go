package main

import (
	"designpatterns/statregy"
	"flag"
	"fmt"
	"log"
)

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {

	flag.Parse()

	var activeStrategy statregy.PrintStrategy

	switch *output {
	case "console":
		activeStrategy = &statregy.ConsoleSquare{}
	case "image":
		activeStrategy = &statregy.ImageSquare{DestinationFilePath: "/tmp/image.jpg"}
	default:
		activeStrategy = &statregy.ConsoleSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}

	s := send()
	r := receive(s)
	fmt.Println(<-r)

}

func send() <-chan int {

	out := make(chan int, 10)

	go func() {

		for i := 0; i < 100000; i++ {
			out <- i
		}
		close(out)
	}()

	return out

}
func receive(in <-chan int) <-chan int {

	out := make(chan int, 10)

	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out

}
