package main

import (
	"designpatterns/statregy"
	"flag"
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

}
