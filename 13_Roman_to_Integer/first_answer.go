package main

import (
	"strings"
)

func main() {
	// romanToInt("III")
	// romanToInt("LVIII")
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	romanMap := map[string]int{}
	roman := []string{}

	romanMap["CM"] = 900
	roman = append(roman, "CM")
	romanMap["CD"] = 400
	roman = append(roman, "CD")
	romanMap["XC"] = 90
	roman = append(roman, "XC")
	romanMap["XL"] = 40
	roman = append(roman, "XL")
	romanMap["IX"] = 9
	roman = append(roman, "IX")
	romanMap["IV"] = 4
	roman = append(roman, "IV")

	romanMap["III"] = 3
	roman = append(roman, "III")
	romanMap["XXX"] = 30
	roman = append(roman, "XXX")
	romanMap["CCC"] = 300
	roman = append(roman, "CCC")
	romanMap["MMM"] = 3000
	roman = append(roman, "MMM")

	romanMap["II"] = 2
	roman = append(roman, "II")
	romanMap["XX"] = 20
	roman = append(roman, "XX")
	romanMap["CC"] = 200
	roman = append(roman, "CC")
	romanMap["MM"] = 2000
	roman = append(roman, "MM")

	romanMap["I"] = 1
	roman = append(roman, "I")
	romanMap["V"] = 5
	roman = append(roman, "V")
	romanMap["X"] = 10
	roman = append(roman, "X")
	romanMap["L"] = 50
	roman = append(roman, "L")
	romanMap["C"] = 100
	roman = append(roman, "C")
	romanMap["D"] = 500
	roman = append(roman, "D")
	romanMap["M"] = 1000
	roman = append(roman, "M")

	a := 0
	before := s
	for _, ro := range roman {
		r := strings.Replace(before, ro, "", 1)
		if r != before {
			a += romanMap[ro]
			before = r
			continue
		}
		if r == "" {
			break
		}
		before = r
	}
	return a
}
