// JavaScript episode N13
// challenge N17 => (Dominant Array Elements) - 7Kyu

// By ALGORITHM

const solve = (arr) => {
  let r = [];
  for (let i = 0; i < arr.length - 1; i++) {
    const element = arr[i];
    const check = (currentNum) => element > currentNum;
    arr.slice(i + 1).every(check) ? r.push(element) : "";
  }
  r.push(arr[arr.length - 1]);
  console.log(r);
};

solve([1, 21, 4, 7, 5]);
solve([5, 4, 3, 2, 1]);
solve([16, 17, 14, 3, 14, 5, 2]);
