package structure

func MaximumLengthSubstring(s string) int {
	var rightIndex int = 0
	var leftIndex int = 0
	var lenTotal int = 1

	var hashMap = map[string]int{string(s[0]): 1}

	for rightIndex < len(s)-1 {
		rightIndex += 1

		if _, existe := hashMap[string(s[rightIndex])]; existe {
			hashMap[string(s[rightIndex])] += 1
		} else {
			hashMap[string(s[rightIndex])] = 1
		}

		for hashMap[string(s[rightIndex])] == 3 {
			hashMap[string(s[leftIndex])] -= 1
			leftIndex += 1
		}

		lenTotal = max(lenTotal, rightIndex-leftIndex+1)
	}

	return lenTotal
}
