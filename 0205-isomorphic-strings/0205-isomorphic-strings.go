func isIsomorphic(s string, t string) bool {
    mapa1 := map[byte]byte{}
    mapa2 := map[byte]byte{}

    for i := 0; i < len(s); i++ {
        if v, ok := mapa1[s[i]]; ok && v != t[i] {
            return false
        }
        if v, ok := mapa2[t[i]]; ok && v != s[i] {
            return false
        }

        mapa1[s[i]] = t[i]
        mapa2[t[i]] = s[i]
    }

    return true
}