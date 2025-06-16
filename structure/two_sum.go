package structure

func twoSum(nums []int, target int) []int {
	var hashMap = map[int]int{}

	for i, num := range nums {
		for idxHash, hashValue := range hashMap {
			if hashValue+num == target {
				nums = []int{idxHash, i}
			}
		}

		hashMap[i] = num
	}

	return nums
}
