package dsa

func search(nums []int, target int) int {
	var rightIndex int = len(nums) - 1
	var leftIndex int = 0

	for leftIndex <= rightIndex {
		var midIndex int = (leftIndex + rightIndex) / 2

		if nums[midIndex] == target {
			return midIndex
		}

		if nums[midIndex] < target {
			leftIndex = midIndex + 1
		}

		if nums[midIndex] > target {
			rightIndex = midIndex - 1
		}
	}

	return -1
}
