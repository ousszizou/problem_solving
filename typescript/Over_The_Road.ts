// TypeScript episode N14
// challenge N18 => (Over The Road) - 7Kyu

// By ALGORITHM ACADEMY

const overTheRoad = (address: number, n: number):number => {
  return (n*2) - address + 1;
}

console.log(overTheRoad(1,3)) // 6
console.log(overTheRoad(3,5)) // 8

/*

overTheRoad(3,5)

5 * 2 = 10

1    10
3    8
5    6
7    4
9    2

overTheRoad(3,5) => 8 : (5*2) - 3 + 1

*/
