/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
   let  myset = new Set(nums)
   return !(nums.length==myset.size)
};