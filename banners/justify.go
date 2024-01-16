package banners

import (
	"strings"

	. "ascii/internal"
)

func GetAsciiJustify(text, data string) string {
	var result string

	var subsubresult [8]string

	table := CreateMap(string(data))

	text = strings.ReplaceAll(text, "\\n", "\n")

	s := CustomSplit(text) // Делим по строкам

	var flag bool

	for i, subs := range s {

		if subs == "\n" {
			if !flag {
				continue
			}
			result += "\n"
			continue
		}

		//
		maxlen, nwords := getMaxlen(subs, table)
		_, cols := GetTerminalSize()

		spaceTokens := cols - maxlen

		spacecosts := distributeSpaceTokens(spaceTokens, nwords-1)
		//

		words := strings.Fields(subs)

		for i, word := range words {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if art, ok := table[char]; ok {
						subsubresult[i] += art[i]
					}
				}
			}

			if i < len(words)-1 {
				for j := 0; j < 8; j++ {
					subsubresult[j] += strings.Repeat(" ", spacecosts[i])
				}
			}

		}

		for i := 0; i < 8; i++ {
			result += subsubresult[i] + "\n"
		}
		subsubresult = [8]string{}

		// checking nextword existing
		if i < len(s)-1 && len(s[i+1]) > 0 {
			flag = true
		}

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

func distributeSpaceTokens(spaceTokens, spaceVauchers int) []int {
	if spaceVauchers == 0 {
		return []int{spaceTokens}
	}

	result := make([]int, spaceVauchers)

	equalShare := spaceTokens / spaceVauchers

	remainder := spaceTokens % spaceVauchers

	for i := 0; i < spaceVauchers; i++ {
		result[i] = equalShare
		if remainder > 0 {
			result[i]++
			remainder--
		}
	}

	return result
}
