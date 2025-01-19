/**
 * Personally I think this one might be nonesense for JS Dev. Because of numbers of string methods to use. Look like cheat, but just a choice.
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
    return x.toString().split('').reverse().join('') === x.toString();
};