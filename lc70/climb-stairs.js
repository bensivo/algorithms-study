function climbStairs(n) {

    // Idea: at each step, compute the number of distinct ways you can get to that step
    // Use that value to calculate the next step

    const results = new Array(n);

    for (let i=1; i<=n; i++) {
        // Pre-set base-cases
        if (i==1) {
            results[i] = 1
            continue;
        }
        if (i==2) {
            results[i] = 2
            continue;
        }

        // Only 2 ways to get to n, from n-2, then take a double-step, or n-1 and take a single step.
        const nPrev = results[i-1]
        const n2Prev = results[i-2]
        results[i] = (nPrev) + (n2Prev)
    }
    return results[n];
}


console.log(climbStairs(1))
console.log(climbStairs(2))
console.log(climbStairs(3))
console.log(climbStairs(4))
console.log(climbStairs(5))
console.log(climbStairs(6))