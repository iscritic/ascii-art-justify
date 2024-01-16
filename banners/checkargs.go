package banners

import (
	"flag"
	"fmt"
)

var validAlignValues = map[string]bool{
	"right":   true,
	"left":    true,
	"center":  true,
	"justify": true,
}

func isValidAlignValue(value string) bool {
	_, valid := validAlignValues[value]
	return valid
}

func CheckArgsJustify() (string, string, string, error) {
	align := flag.String("align", "", "Specify alignment (right, left, center, justify)")

	flag.Parse()

	if *align != "" && !isValidAlignValue(*align) {
		return "", "", "", fmt.Errorf("Invalid align value: %s", *align)
	}

	args := flag.Args()

	if len(args) == 1 {
		return *align, args[0], "standard", nil
	}

	if len(args) == 2 {
		return *align, args[0], args[1], nil
	}

	// Возвращаем значения
	return "", "", "", nil
}
