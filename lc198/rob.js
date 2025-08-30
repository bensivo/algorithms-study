
function rob(arr) {
    // Base case, 0 length arr
    if (arr.length == 0) {
        return 0;
    }

    // Saves the max money you can get for N number of houses
    const results = new Array(arr.length);

    for(let i=0; i<arr.length; i++) {
        // Base cases, 1 or 2 items
        if (i==0) {
            results[i] = arr[i];
            continue;
        }
        if (i==1) {
            results[i] = Math.max(arr[0], arr[1]);
            continue;
        }

        // Money you'd get by skipping this house (same as if this were not there)
        const skip = results[i-1];

        // Money you'd get by skipping the previous house (same as if prev were not there, but you added this houses value)
        const take = results[i-2] + arr[i];

        results[i] = take > skip ? take : skip;
    }

    return results[arr.length -1]
}


// WOrking it out by hand
/**
[1,2,3,1]

[1]: 1 = pick 1
[1,2,]: 2 = rob new is better, pick 2 (n-1 + value of new)
[1,2,3]: 4 = rob new is better, pick 3 (n-2 + value of new)
[1,2,3,1]: 4 = n-1 = 4, n-2 + new = 3, keep n-1,
 */

console.log(rob([1,]))
console.log(rob([1,2,]))
console.log(rob([1,2,3]))
console.log(rob([1,2,3,1]))
console.log(rob([1,2,3,1, 7]))
console.log(rob([2,7,9,3,1]))