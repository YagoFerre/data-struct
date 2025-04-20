package dsa

import (
	"fmt"
)

func main() {
	testCases := "Let's take LeetCode contest"

	fmt.Printf("Original: %q\n", testCases)
	fmt.Printf("Reversed: %q\n\n", reverseWord(testCases))
}
