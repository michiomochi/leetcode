package main

func longestCommonPrefix(strs []string) string {
	a := strs[0]

	for _, str := range strs[1:] {
		aa := ""
		
		for i, aw := range a {
			if (len(str) < i + 1) {
				break
			}

			strw := rune(str[i])

			if (aw == strw) {
				aa += string(aw)
			} else {
				break
			}
		}

		a = aa

		if (a == "") {
			break
		}
	}

	return a
}