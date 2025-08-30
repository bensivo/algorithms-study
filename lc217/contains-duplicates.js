function lc217(arr) {

    valueMap = {} // maps values to boolean

    for(let i=0; i<arr.length; i++) {
        const value = arr[i];
        if (valueMap[value]) {
            return true;
        }

        valueMap[value] = true;
    }
    return false;
}

console.log(lc217([1,2,3,1]))

console.log(lc217([1,2,3,4]))

console.log(lc217([1,1,1,3,3,4,3,2,4,2]))