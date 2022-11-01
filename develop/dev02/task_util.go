package develop

import "strings"

func isNumeric(char rune) bool {
	return strings.ContainsRune(numeric, char)
}

func isEscapable(char rune) bool {
	return strings.ContainsRune(escapable, char)
}

func parseNum(str string) (num string) {
	i := 0
	for i < len(str) && isNumeric(rune(str[i])) {
		i++
	}

	return str[:i]
}
