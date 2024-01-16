package main

import (
	"fmt"
	"log"
	"os"

	"ascii/banners"
	"ascii/internal"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
		return
	}

	align, text, font, err := banners.CheckArgsJustify()
	if err != nil {
		fmt.Println(err)
		return
	}

	if isValidText := internal.IsValidText(text); !isValidText {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard	")
		return
	}

	if text == "" {
		log.Println("Text argument is empty")
		return
	}

	data, err := internal.FontPicker(font)
	if err != nil {
		log.Println(err)
		return
	}

	output := banners.AsciiJustify(text, data, align)

	fmt.Print(output)
}
