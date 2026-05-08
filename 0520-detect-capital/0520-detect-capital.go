func detectCapitalUse(word string) bool {
low := strings.ToLower(word)

if strings.ToUpper(string(low[0]))+low[1:] == word || low == word || strings.ToUpper(low) == word {
    return true
}
return false
}