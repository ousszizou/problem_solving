// JavaScript episode N15
// challenge N19 => (Shifter words) - 7Kyu

// By ALGORITHM ACADEMY

/*
  We will call a word "shifter" if it consists only of letters "H", "I", "N", "O", "S", "X", "Z", "M" and "W"
*/

// SOS IN THE HOME

const letters = ["H", "I", "N", "O", "S", "X", "Z", "M", "W"];

const shifter = (s) => {
  let sum = 0;
  if (s.length == 0) {
    return 0;
  }
  const unique_words = [...new Set(s.split(" "))];
  unique_words.map((word) => {
    word.split("").every((l) => letters.includes(l)) ? (sum += 1) : "";
  });
  return sum;
};

console.log(shifter("SOS IN THE HOME")); // 2
console.log(shifter("WHO IS SHIFTER AND WHO IS NO")); // 3
console.log(shifter("TASK")); // 0
console.log(shifter("")); //0
