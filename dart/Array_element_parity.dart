// Dart episode N08
// challenge N14 = > (Array element parity) - 7Kyu

// By ALGORITHM ACADEMY


// ignore: missing_return
int solve(List arr) {
  for (var n in arr) {
    if (arr.contains(-n) == true) {
      continue;
    } else {
      return n;
    }
  }
}

main() {
  print(solve([1, -1, 2, -2, 3]));
  print(solve([-3, 1, 2, 3, -1, -4, -2]));
}
