package structure

func FirstUniqueChar(s string) int {
	var hashMap = map[string][]int{}

	for idx, ch := range s {
		if _, existe := hashMap[string(ch)]; !existe {
			hashMap[string(ch)] = []int{idx, 1}
		} else {
			hashMap[string(ch)][1] += 1
		}
	}

	// idx _
	for _, hashValue := range hashMap {
		if hashValue[1] == 1 {
			return hashValue[0]
		}
	}

	return -1
}
