function lc1(arr, target) {
    const valueIndexMap = {} // Maps the value to the index of that value

    for (let i=0; i<arr.length; i++) {
        const diff = target - arr[i];

        // see if needed value has already been seen
        if (valueIndexMap[diff] !== undefined) {
            return [valueIndexMap[diff], i]
        }

        valueIndexMap[arr[i]] = i;
    }
}



console.log(lc1([2,7,11,15], 9));

console.log(lc1([3,2,4], 6));

console.log(lc1([3,3], 6));