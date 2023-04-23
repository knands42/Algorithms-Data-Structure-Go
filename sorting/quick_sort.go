package sorting

import "math/rand"

// QuickSort Best case: O(n log n)
// Average case: O(n log n)
// Worst case: O(n^2)
func QuickSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength < 2 {
		return arr
	}

	idx := getPivotIdx(arrLength)
	pivot := arr[idx]
	var left, right []int

	for i := 0; i < arrLength; i++ {
		if i == idx {
			continue
		}
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

func getPivotIdx(arrLength int) int {
	return rand.Intn(arrLength)
}
