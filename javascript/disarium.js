// JavaScript episode N02
// challenge N06 => (Disarium Number) - 7Kyu

// By ALGORITHM


function disariumNumber(n) {
  const arr = Array.from(String(n), Number)
  let sum = 0
  for (let index = 0; index < arr.length; index++) {
    sum += arr[index] ** (index + 1)
  }
  return sum == n ? 'Disarium !!' : 'Not !!' 
}

console.log(disariumNumber(89))
console.log(disariumNumber(189))