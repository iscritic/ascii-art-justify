package banners

import (
	"fmt"
	"log"
	"strings"

	. "ascii/internal"
)

func GetAsciiJustify(text, data string) string {
	var result string

	var subsubresult [8]string

	table := CreateMap(string(data))

	text = strings.ReplaceAll(text, "\\n", "\n")

	s := CustomSplit(text) // Делим по строкам

	// for _, str := range s {
	// 	fmt.Println("!" + str + "!")
	// }

	for _, subs := range s {

		// if subs == "\n" {
		// 	if !flag {
		// 		continue
		// 	}
		// 	result += "\n"
		// 	continue
		// }

		//
		maxlen, nwords := getMaxlen(subs, table)
		_, cols := GetTerminalSize()

		spaceTokens := cols - maxlen

		spacecosts := distributeCandies(spaceTokens, nwords-1)
		//

		fmt.Println(spacecosts)

		words := strings.Fields(subs)

		for i, word := range words {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if art, ok := table[char]; ok {
						subsubresult[i] += art[i]
						// fmt.Println(subsubresult[i])
					}
				}
			}

			// for i := 0; i < 8; i++ {
			// 	subsubresult[i] += strings.Repeat(" ", spacecosts[j])
			// }
			if i == 0 {
				subsubresult[0] += strings.Repeat(" ", 102)
				subsubresult[1] += strings.Repeat(" ", 102)
				subsubresult[2] += strings.Repeat(" ", 102)
				subsubresult[3] += strings.Repeat(" ", 102)
				subsubresult[4] += strings.Repeat(" ", 102)
				subsubresult[5] += strings.Repeat(" ", 102)
				subsubresult[6] += strings.Repeat(" ", 102)
				subsubresult[7] += strings.Repeat(" ", 102)
			}

			if i == 1 {
				subsubresult[0] += strings.Repeat(" ", 101)
				subsubresult[1] += strings.Repeat(" ", 101)
				subsubresult[2] += strings.Repeat(" ", 101)
				subsubresult[3] += strings.Repeat(" ", 101)
				subsubresult[4] += strings.Repeat(" ", 101)
				subsubresult[5] += strings.Repeat(" ", 101)
				subsubresult[6] += strings.Repeat(" ", 101)
				subsubresult[7] += strings.Repeat(" ", 101)
			}
		}

		for i := 0; i < 8; i++ {
			fmt.Println(subsubresult[i])
		}

		log.Fatalln()
		// result += subresult + "\n"
		// subresult = ""

		// // checking nextword existing
		// if i < len(s)-1 && len(s[i+1]) > 0 {
		// 	flag = true
		// }

	}

	return result
}

func getMaxlen(subs string, m map[rune][]string) (int, int) {
	var tmp string
	words := strings.Fields(subs)

	for _, word := range words {
		for _, char := range word {
			tmp += string(m[char][0])
		}
	}

	return len(tmp), len(words)
}

func distributeCandies(candies, people int) []int {
	result := make([]int, people)

	// Равное количество конфет для каждого
	equalShare := candies / people

	// Оставшееся количество конфет, которое нужно распределить
	remainder := candies % people

	for i := 0; i < people; i++ {
		result[i] = equalShare
		if remainder > 0 {
			result[i]++
			remainder--
		}
	}

	return result
}
