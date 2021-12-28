// TypeScript episode N16
// challenge N20 => (Array.diff) - 6Kyu

// By ALGORITHM ACADEMY

function arrayDiff(a: number[], b: number[]): number[] {
  return a.filter((n) => !b.includes(n));
}

console.log(arrayDiff([], [4, 5])); // []
console.log(arrayDiff([3, 4], [3])); // [4]
console.log(arrayDiff([1, 8, 2], [])); // [1,8,2]
console.log(arrayDiff([1, 2, 3], [1, 2])); // [3]
