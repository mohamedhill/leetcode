func reverseVowels(s string) string {
	r := []rune(s)

	i := 0
    j := len(r)-1
  

	for i < j {
		if !isvowel(r[i]) {
			i++
			continue
		}
		if !isvowel(r[j]) {
			j--
			continue
		}

		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}

func isvowel(vowel rune) bool {
	vowels := []rune{'a','e','i','o','u','A','E','I','O','U'}

	for _, v := range vowels {
		if vowel == v {
			return true
		}
	}
	return false
}