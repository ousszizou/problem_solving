// JavaScript episode N17
// challenge N21 => (Negative Connotation) - 7Kyu

// By ALGORITHM ACADEMY

const connotation = (str) => {
  let countPositiveWords = 0;
  let countNegativeWords = 0;

  const words = str.toLowerCase().trim().replace(/\s+/g, " ").split(" ");

  words.map((w) => {
    if (w.charCodeAt(0) <= 109) {
      countPositiveWords++;
    } else {
      countNegativeWords++;
    }
  });

  return countPositiveWords >= countNegativeWords ? true : false;
};

console.log(connotation("A big brown fox caught a bad bunny")); // true
console.log(connotation("Xylophones can obtain Xenon.")); // false
console.log(connotation("Is  this the  best   Kata?")); // true
console.log(connotation("All FOoD tAsTEs NIcE for someONe")); // true
