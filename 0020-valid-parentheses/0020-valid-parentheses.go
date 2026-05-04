func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
    
	p := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stc := []rune{}

	for _, r := range s {
		if _, ok := p[r]; ok {
			stc = append(stc, r)
		} else if len(stc) == 0 || p[stc[len(stc)-1]] != r {
			return false
		} else {
			stc = stc[:len(stc)-1]
		}
	}

	return len(stc) == 0
}