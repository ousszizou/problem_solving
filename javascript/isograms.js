// JavaScript episode N02
// challenge N04 => (Isograms) - 7Kyu

// By ALGORITHM


function isIsogram(str) {
  return new Set(str.toLowerCase()).size == str.length
}

console.log(isIsogram('azerty'))
console.log(isIsogram('NBNA'))