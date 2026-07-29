func plusOne(digits []int) []int {
    nine := -1

    if digits[len(digits)-1]==9{
   for i := len(digits)-1;i>=0;i--{
    if digits[i]!= 9 {
        
        break
    }
    nine = i

   }
    if nine == 0{
        digits[nine] = 1
        digits = append(digits,0)
        for i := nine+1 ; i <len(digits);i++{
            digits[i]= 0
        }

    }else {
      digits[nine-1]= digits[nine-1]+1
     for i := nine ; i <len(digits);i++{
            digits[i]= 0
        }
    }



    }else {
    digits[len(digits)-1] = digits[len(digits)-1]+1

   }
 
     return digits 
    
}
