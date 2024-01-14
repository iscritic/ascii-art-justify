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

		spaceTokens := cols - 68 - maxlen

		spacecosts := distributeCandies(spaceTokens, nwords-1)
		//

		words := strings.Fields(subs)

		var j int

		for _, word := range words {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if art, ok := table[char]; ok {
						subsubresult[i] += art[i]
						// fmt.Println(subsubresult[i])
					}
				}
				subsubresult[i] += strings.Repeat(" ", spacecosts[j])
				if j < len(spacecosts)-1 {
					j++
				}

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
