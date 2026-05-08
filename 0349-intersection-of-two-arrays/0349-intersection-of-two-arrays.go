func intersection(nums1 []int, nums2 []int) []int {
maps := make(map[int]bool)
 var result [] int 
for _, i := range nums1{
maps[i]=true 
}

for _,j := range nums2{
    if maps[j]==true{
        result = append(result,j)
        maps[j]=false
    }
}
return result
}