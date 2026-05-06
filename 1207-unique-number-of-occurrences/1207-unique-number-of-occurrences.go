func uniqueOccurrences(arr []int) bool {
    mapa := make(map[int]int)


    for _, j := range arr {
        mapa[j]++
    }

    
    seen := make(map[int]bool)

    for _, count := range mapa {
        if seen[count] {
            return false 
        }
        seen[count] = true
    }

    return true
}