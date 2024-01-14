package main

import (
	"fmt"
	"log"
	"os"

	. "ascii/banners"
	. "ascii/internal"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("At least 2 arguments are required")
	}
	arg := os.Args[1]

	if arg == "" {
		return
	}

	data, err := FontPicker()
	if err != nil {
		log.Fatalln("Invalid font")
	}

	output := GetAsciiJustify(arg, data)

	fmt.Print(output)
}
