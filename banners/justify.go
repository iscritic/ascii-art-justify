package banners

import (
	"strings"

	. "ascii/internal"
)

func GetAsciiJustify(text, data string) (string, int) {
	table := CreateMap(string(data))

	text = strings.ReplaceAll(text, "\\n", "\n")

	s := CustomSplit(text)

	var result, subresult string
	var maxlen int
	var flag bool

	for i := 0; i < len(s); i++ {

		subs := s[i]

		if subs == "\n" {
			if flag {
				flag = false
				continue
			}
			result += "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range subs {
				if art, ok := table[char]; ok {
					subresult += art[i]
				}
			}

			// Get len of longest string
			if maxlen < len(subresult) {
				maxlen = len(subresult)
			}

			result += subresult + "\n"
			subresult = ""

		}

		// checking nextword existing
		if i < len(s)-1 && len(s[i+1]) > 0 {
			flag = true
		}
	}

	return result, maxlen
}
