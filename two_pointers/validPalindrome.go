package easy

import (
	"regexp"
	"strings"
)

func sanitizeText(s string) string {
	onlyLettersText := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	return strings.ToLower(onlyLettersText)
}

func IsPalindrome(text string) bool {
	text = sanitizeText(text)
	textSize := len(text)

	if textSize == 0 {
		return true
	}

	left := 0
	right := textSize - 1

	for left < right {
		if text[left] != text[right] {
			return false
		}

		left++
		right--
	}

	return true
}
