// JavaScript episode N08
// challenge N14 = > (Array element parity) - 7Kyu

// By ALGORITHM ACADEMY

function solve(arr) {
  for (let i = 0; i < arr.length; i++) {
    const n = arr[i];
    r = arr.filter(x => x == -n).length
    
    if (r > 0) {
      continue
    } else {
      return n
    }
  }
}

console.log(solve([1, -1, 2, -2, 3]))
console.log(solve([-3, 1, 2, 3, -1, -4, -2]))
