// https://leetcode.com/problems/valid-anagram/
package arrays

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[string]int)

	for _, v := range s {
		dict[string(v)] += 1
	}

	for _, char := range t {
		char := string(char)
		if amount, ok := dict[char]; ok {
			if amount == 1 {
				delete(dict, char)
			} else {
				dict[char] -= 1
			}
		} else {
			return false
		}
	}

	return len(dict) == 0
}
