package sorting

// MergeSort Best case: O(n log n)
// Average case: O(n log n)
// Worst case: O(n log n)
func MergeSort(arr []int) []int {
	arrLength := len(arr)
	if arrLength < 2 {
		return arr
	}

	mid := arrLength / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))

	for len(arr1) > 0 || len(arr2) > 0 {
		if len(arr1) == 0 {
			return append(result, arr2...)
		}

		if len(arr2) == 0 {
			return append(result, arr1...)
		}

		if arr1[0] <= arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}

	return result
}
