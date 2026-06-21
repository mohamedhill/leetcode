
func maxIceCream(costs []int, coins int) int {
    costs = countingSort(costs)
    sums := 0
    test :=0

    for i := 0 ; i < len(costs);i++{
    

            sums+= costs[i]
        if sums<=coins{
            test++
        
        }else if i != 0 {
            sums -= costs[i]
        }

    }
   
    return test
    
}

func countingSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)

	for _, v := range arr {
		count[v]++
	}

	result := []int{}

	for value, freq := range count {

		for i := 0; i < freq; i++ {
			result = append(result, value)
		}
	}

	return result
}