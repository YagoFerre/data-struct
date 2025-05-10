package structure

func ReverseWord(s string) string {
	var result string = ""
	var reverseIndex int = -1

	// len(s) - 1 ultimo valor do array index
	for rightIndex := 0; rightIndex < len(s); rightIndex++ { // 1,2,3,4,5,6
		if rightIndex == len(s)-1 || s[rightIndex] == ' ' {
			// diminuir o array de string
			var leftIndex int

			if rightIndex == len(s)-1 {
				leftIndex = rightIndex // ultimo index da string
			}

			if s[rightIndex] == ' ' {
				leftIndex = rightIndex - 1 // ultimo index da string se tiver espaço
			}

			// jogar o leftIndex para prox posiçao da string
			// atribuir valor ao result
			for _ = 0; leftIndex > reverseIndex; leftIndex-- {
				result += string(s[leftIndex])
			}

			if rightIndex != len(s)-1 {
				result += " " // adiciona espaço entre palavras
			}

			reverseIndex = rightIndex
		}
	}

	return result
}
