package structure

func ExponentialSearch(nums []int, target int) int {
	// se for o primeiro
	if nums[0] == target {
		return 0
	}

	var arrSize int = len(nums)
	var rightIndex int = 1

	// dobra a busca pelo index
	for rightIndex < arrSize && nums[rightIndex] < target {
		rightIndex *= 2
	}

	startSubArr := rightIndex / 2

	if rightIndex >= arrSize {
		rightIndex = arrSize - 1
	}

	endSubArr := rightIndex

	// se for o ultimo
	if nums[rightIndex] == target {
		return rightIndex
	}

	return Search(nums[startSubArr:endSubArr], target)
}
