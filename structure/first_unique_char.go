package structure

func FirstUniqueChar(s string) int {
	var hashMap = map[string][]int{}

	// hashMap: {l: [0, 1], e: [1, 2], t: [3, 3]}
	for idx, ch := range s {
		if _, existe := hashMap[string(ch)]; !existe {
			hashMap[string(ch)] = []int{idx, 1}
		} else {
			hashMap[string(ch)][1] += 1
		}
	}

	// percorre string original com hashMap montado
	for idx, ch := range s {
		if hashMap[string(ch)][1] == 1 {
			return idx
		}
	}

	return -1
}
