package structure

func containsNearbyDuplicate(nums []int, k int) bool {
	var hashMap = map[int]int{} // [{123123:40}, {234243:76}, {567564:80}] nao duplica so muda o valor da chave

	for i, num := range nums {
		// [{123123:40}] --> 40 key hashMap !-- index arr
		if lastIndexValueHashMap, exists := hashMap[num]; exists {
			if i-lastIndexValueHashMap <= k {
				return true
			}
		}

		// [{123123:47}] sobreescreve o valor da chave existente
		hashMap[num] = i
	}

	return false
}
