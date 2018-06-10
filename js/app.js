function multiplyBetween(num1, num2) {
    if (isNaN(num1) || isNaN(num2)) return NaN;
    if (num1 === 0 || num2 === 0) return 0;
    if (num1 === num2) return 1;
    if (num2 < num1) {
        return 0;
    }
    return num1 * multiplyBetween(num1 + 1, num2)
}

console.log(multiplyBetween(3141592653589793238462643383279502884197169399375105820974944592, 2718281828459045235360287471352662497757247093699959574966967627));