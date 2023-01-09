package easy

type Stack []string

func (s *Stack) Push(item string) {
	*s = append((*s), item)
}
func (s *Stack) Pop() (string, bool) {
	if s.Empty() {
		return "", true
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	if s.Empty() {
		s.Clear()
		return element, true
	}
	return element, false
}
func (s *Stack) Peek() (string, bool) {
	if s.Empty() {
		return "", true
	}

	element := (*s)[len(*s)-1]
	return element, false
}
func (s *Stack) Empty() bool {
	return len(*s) < 1
}
func (s *Stack) Length() int {
	return len(*s)
}
func (s *Stack) Clear() {
	*s = []string{}
}

func IsValid(s string, stack Stack) bool {
	closeToOpen := make(map[string]string)
	closeToOpen[")"] = "("
	closeToOpen["]"] = "["
	closeToOpen["}"] = "{"

	for _, char := range s {
		charString := string(char)
		if _, ok := closeToOpen[charString]; !ok {
			stack.Push(charString)
		} else {
			stackValue, empty := stack.Peek()
			if !empty && stackValue == closeToOpen[charString] {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	return stack.Empty()
}
