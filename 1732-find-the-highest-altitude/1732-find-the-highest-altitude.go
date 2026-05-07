func largestAltitude(gain []int) int {
    alt := 0
    max := 0

    for _, g := range gain {
        alt += g

        if alt > max {
            max = alt
        }
    }

    return max
}