package main

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}
	for _, c := range s {
		if _, ok := m[c]; ok {
			stack = push(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			var key rune
			stack, key = pop(stack)

			if v, ok := m[key]; ok {
				if v != c {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func push(l []rune, v rune) []rune {
	return append(l, v)
}

func pop(l []rune) ([]rune, rune) {
	v := l[len(l)-1]
	l2 := l[:len(l)-1]
	return l2, v
}
