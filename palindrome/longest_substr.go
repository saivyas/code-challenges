package main

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	lStart, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}
		// skip duplicates
		k, j := i, i
		for k < len(s)-1 && s[k] == s[k+1] {
			k++
		}
		// expand around centre
		i = k + 1
		for k < len(s)-1 && j > 0 && s[k+1] == s[j-1] {
			k++
			j--
		}
		newLen := k - j + 1
		if newLen > maxLen {
			maxLen = newLen
			lStart = j
		}
	}
	return s[lStart : lStart+maxLen]
}
func reverseStr(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	longestPalindrome("babad")
}
