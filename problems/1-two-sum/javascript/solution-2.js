/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    let obj = {}
    for(let i = 0, n = nums.length; i < n; i++) {
        if(obj[target - nums[i]] !== undefined) {
            return [obj[target - nums[i]], i];
        }
        obj[nums[i]] = i;
    }
    return [];
};