func isPalindrome(x int) bool {
    sa := []rune(strconv.Itoa(x))
    for i := 0; i < len(sa) / 2; i++ {
        if (sa[i] != sa[len(sa) -i - 1]) {
            return false
        }
    }
    return true
}