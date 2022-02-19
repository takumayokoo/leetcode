package main

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	if len(strs) == 1 {
		return prefix
	}

	for _, s := range strs[1:] {
		length := 0

		if len(prefix) < len(s) {
			length = len(prefix)
		} else {
			length = len(s)
		}

		newPrefix := ""

		for i := 1; i <= length; i++ {
			if prefix[0:i] == s[0:i] {
				newPrefix = prefix[0:i]
			}
		}

		prefix = newPrefix
	}

	return prefix
}
