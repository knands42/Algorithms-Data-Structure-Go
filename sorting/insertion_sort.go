package sorting


func insertionSort(numbers []int) []int {
  for key := 1; key < len(numbers); key++ {
    currentKey := key
    for currentKey > 0 && numbers[currentKey-1] > numbers[currentKey] {
      numbers[currentKey], numbers[currentKey - 1] = numbers[currentKey - 1], numbers[currentKey]
      currentKey = currentKey - 1
    }
  }
  return numbers
}
