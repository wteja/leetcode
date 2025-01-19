/**
 * This one also be the good choice, similar to solution 1, but don't need to compare full string.
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    const arr = String(x).split('');
    while(arr.length > 1) {
        if(arr.shift() !== arr.pop()) {
            return false;
        }
    }
    return true;
};