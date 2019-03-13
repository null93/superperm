const { alphabet, brute_force } = require ("../../src")

let permutations = alphabet ( 3 )
let result = brute_force ( permutations )
console.log ( result.length, result )
