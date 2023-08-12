package main

import "github.com/k0kubun/pp/v3"

func main() {
	isValid("()[]{}")
}

func isValid(s string) bool {
	pare := map[string]string{
		"]": "[",	
		")": "(",
		"}": "{",
	}
	stack := []string{}

	for _, ss := range s {
		sss := string(ss)
		if ssv, exists := pare[sss]; exists {
			if len(stack) > 0 {
				if stack[len(stack)-1] == ssv {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, sss)
		}
	}

	if len(stack) > 0 {
		return false
	}
	
	return true
}
