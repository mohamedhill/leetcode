impl Solution {
    pub fn left_right_difference(nums: Vec<i32>) -> Vec<i32> {
        let mut leftsum = 0 ;
        let mut rightsum = 0;

        let mut leftvec = vec![];
        let mut rightvec = vec![];
        let mut res = vec![];

        for i in &nums{
            leftvec.push(leftsum);
            leftsum+=i;

        };
        for m in (0..nums.len()).rev(){
            rightvec.push(rightsum);
            rightsum+=nums[m];

        };
        rightvec.reverse();
let mut index :usize =0;
for k in leftvec{
    if k > rightvec[index]{
        res.push(k-rightvec[index]);
    }else{
        res.push(rightvec[index]-k);

    };
    index+=1
};

res
    }
   
}