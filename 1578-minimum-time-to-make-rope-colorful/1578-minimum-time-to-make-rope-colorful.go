func minCost(colors string, neededTime []int) int {
    current := ""
	res := 0
    for i := 0 ; i<len(colors);i++{
        if current == string(colors[i]){
			if neededTime[i-1] > neededTime[i]{
				
			res += neededTime[i]
				neededTime[i] = neededTime[i-1]
		}else{
			res += neededTime[i-1]
		}
	
        }
        current = string(colors[i])

    }
   
	return res
}