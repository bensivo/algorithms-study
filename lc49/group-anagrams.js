function lc49(arr) {

    const groups = {} // Maps anagram key to list of strings in that group

    // For string
    for (let i=0; i<arr.length; i++) {
        // sort letters alphabetically, create lookup key

        const str = arr[i];
        const key = sortStr(str);

        if (!groups[key]) {
            groups[key] = [str];
        } else {
            groups[key].push(str);
        }
    }

    return Array.from(Object.entries(groups).map(([k,v]) => v))
}

// Alternate version, just count occurences of each of the 26 characters
function sortStr(str) {
    const cstr = new Array(str.length)
    for(let i=0; i<str.length; i++) {
        cstr[i] = str[i]
    }
    cstr.sort();
    return cstr;
}


console.log(lc49(["eat","tea","tan","ate","nat","bat"]))
console.log(lc49([""]))
console.log(lc49(["a"]))