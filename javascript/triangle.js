// JavaScript episode N10
// challenge N15 => (Which triangle is that?) - 7Kyu

// By ALGORITHM

// Kata's url : https://www.codewars.com/kata/564d398e2ecf66cec00000a9

var typeOfTriangle = function (sideA, sideB, sideC) {
  if (
    typeof sideA === typeof sideB &&
    typeof sideA === typeof sideB &&
    typeof sideA === "number"
  ) {
    if (sideA == sideB && sideA == sideC && sideA != null) {
      return "Equilateral"
    }
    if (sideA == sideB || sideA == sideC || sideB == sideC) {
      return "Isosceles"
    }
    if (sideA != sideB && sideA != sideC) {
      return "Scalene"
    }
  } else {
    return "Not a valid triangle"
  }
};

console.log(typeOfTriangle(1, 1, 1)); // Equilateral
console.log(typeOfTriangle(2, 2, 1)); // Isosceles
console.log(typeOfTriangle(3, 2, 4)); // Scalene
console.log(typeOfTriangle(".", 5, 8)); // Not a valid triangle
