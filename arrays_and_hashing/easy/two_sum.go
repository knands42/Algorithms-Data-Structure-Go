// https://leetcode.com/problems/two-sum/
package arrays

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		if v, ok := dict[diff]; ok {
			return []int{v, i}
		}

		dict[v] = i
	}

	return []int{}
}
