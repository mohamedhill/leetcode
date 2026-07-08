
func countDays(days int, meetings [][]int) int {
    slices.SortFunc(meetings, func(a, b []int) int {
        return a[0] - b[0]
    })

    busy := 0

    start := meetings[0][0]
    end := meetings[0][1]

    for i := 1; i < len(meetings); i++ {
        if meetings[i][0] <= end {
            if meetings[i][1] > end {
                end = meetings[i][1]
            }
        } else {
            busy += end - start + 1
            start = meetings[i][0]
            end = meetings[i][1]
        }
    }

    busy += end - start + 1

    return days - busy
}
