package rws

func reverseWord(s string) string {
	var result string = ""
	lastStrIndex := -1

	for strIndex := 0; strIndex < len(s); strIndex++ {
		var lastIndexPosition int // last element of string

		if strIndex == len(s)-1 || s[strIndex] == ' ' {
			if strIndex == len(s)-1 {
				lastIndexPosition = strIndex
			} else {
				lastIndexPosition = strIndex - 1
			}

			for _ = 0; lastIndexPosition > lastStrIndex; lastIndexPosition-- {
				result += string(s[lastIndexPosition])
			}

			if strIndex != len(s)-1 {
				result += " "
			}

			lastStrIndex = strIndex
		}
	}

	return result
}
