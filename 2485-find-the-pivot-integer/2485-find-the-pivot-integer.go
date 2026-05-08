func pivotInteger(n int) int {
   
   
        sum:=0
        for i :=1;i<=n;i++{
            sum+=i
        }
        sum2:=0
        for i :=1;i<=n;i++{
            sum2+=i
            if sum==sum2{
                return i
            }
            sum-=i
        }

        return -1
    

}