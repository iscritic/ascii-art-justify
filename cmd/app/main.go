package main

import (
	"fmt"
	"log"
	"os"

	. "ascii/banners"
	. "ascii/internal"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		log.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	align, text, font, err := CheckArgsJustify()
	if err != nil {
		log.Println(err)
		return
	}

	if text == "" {
		fmt.Println("Text argument is empty")
		return
	}

	data, err := FontPicker(font)
	if err != nil {
		log.Println(err)
		return
	}

	output := AsciiJustify(text, data, align)

	fmt.Print(output)
}
