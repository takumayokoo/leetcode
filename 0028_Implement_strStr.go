package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	index := -1

	for i, r := range haystack {
		if rune(needle[0]) == r {
			if len(haystack[i:]) >= len(needle) {
				if haystack[i:i + len(needle)] == needle {
					index = i
					break
				}
			} else {
				break
			}
		}
	}

	return index
}