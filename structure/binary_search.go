package structure

func Search(nums []int, target int, leftIndex int, rightIndex int) int {
	for leftIndex <= rightIndex {
		var midIndex int = (leftIndex + rightIndex) / 2

		if nums[midIndex] == target {
			return midIndex
		}

		if nums[midIndex] < target {
			leftIndex = midIndex + 1 // descartando ate mesmo o mid index
		}

		if nums[midIndex] > target {
			rightIndex = midIndex - 1
		}
	}

	return -1
}
