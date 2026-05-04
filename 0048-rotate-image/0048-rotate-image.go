func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
	matrix[i]= reverse(matrix[i])
	}
	

}


func reverse(arr[]int)[]int{
	res := []int{}
	for i := len(arr)-1 ; i >=0;i--{
		res = append(res, arr[i])
		
	}
	return res
}