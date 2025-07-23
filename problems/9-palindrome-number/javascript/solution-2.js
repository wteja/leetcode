/**
 * I personally love this algorithm, we made up new value each digit by multiply each digit by 10.
 * Then compare it later to see if they have the same value.
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    if(x < 0)
        return false;
    
    let ret = 0;
    for(let i = x; i >= 1; i = Math.floor(i / 10))
        ret = ret * 10 + i % 10;
            
    return ret === x;
};