package LeetCode

import(
	"strconv"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    s := strconv.Itoa(x)

    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }

    return true
}

// Forma mais sofisticada de resolver


func isPalindromeMelhorado(x int) bool {
    // Negativos nunca são palíndromos
    // Números que terminam em 0 (exceto o próprio 0) também não
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }

    reversed := 0
    for x > reversed {
        reversed = reversed*10 + x%10
        x /= 10
    }

    // x == reversed → número par de dígitos (ex: 1221)
    // x == reversed/10 → número ímpar de dígitos (ex: 121, ignoramos o dígito do meio)
    return x == reversed || x == reversed/10
}