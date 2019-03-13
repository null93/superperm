const { alphabet, cycles } = require ("../../src")

let permutations = alphabet ( 4 )
let result = cycles ( permutations )
console.log ( result.length, result )
