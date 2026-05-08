func firstUniqChar(s string) int {
    var uniq =make(map[rune]int)
    for _,v:=range s {
        uniq[v]++
    }
    for i,v:=range s {
        if uniq[v]==1{
            return i
        }
    }
    return -1
}