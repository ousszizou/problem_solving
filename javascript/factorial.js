// JavaScript episode N02
// challenge N05 => (Factorial) - 7Kyu

// By ALGORITHM


function factorial(n) {
  if (n == 0) {
    return 1
  } else if (n == 1) {
    return 1
  } else {
    return n * factorial(n - 1)
  }
}

console.log(factorial(0))
console.log(factorial(1))
console.log(factorial(4))