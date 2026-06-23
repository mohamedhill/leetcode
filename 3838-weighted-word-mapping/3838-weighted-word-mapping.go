func mapWordWeights(words []string, weights []int) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	sums := ""

	for i := 0; i < len(words); i++ {
		curr := 0
		for j := 0; j < len(words[i]); j++ {
			curr += weights[(rune(words[i][j] - 'a'))]
		}
		sums += string(alpha[len(alpha)-(curr%26)-1])
	}
	return sums
}