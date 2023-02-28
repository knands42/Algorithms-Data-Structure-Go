package sorting

func insertionSort(numbers []int) []int {
	for key := 1; key < len(numbers); key++ {
		currentKey := key
		for currentKey > 0 && numbers[currentKey-1] > numbers[currentKey] {
			numbers[currentKey], numbers[currentKey-1] = numbers[currentKey-1], numbers[currentKey]
			currentKey = currentKey - 1
		}
	}
	return numbers
}

// Solution borrowed code from the book "Introduction to Algorithms"
// by Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, and Clifford Stein
func insertionSort2(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		previousIndex := i - 1
		for previousIndex >= 0 && numbers[previousIndex] > number {
			numbers[previousIndex+1] = numbers[previousIndex]
			previousIndex = previousIndex - 1
		}
		numbers[previousIndex+1] = number
	}
	return numbers
}
