package leetcode

// TwoSum indices of the two numbers such that they add up to target.
//
// Time complexity: O(n)
//
// Space complexity: O(n)
func TwoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for currIdx, currNum := range nums {
		if requiredIdx, isPresent := hashmap[target-currNum]; isPresent {
			return []int{requiredIdx, currIdx}
		}
		hashmap[currNum] = currIdx
	}
	return []int{}
}
