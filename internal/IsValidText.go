package internal

func IsValidText(text string) bool {
	for _, ch := range text {
		if !(ch >= 32 && ch <= 127) {
			return false
		}
	}
	return true
}
