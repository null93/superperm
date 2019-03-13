const { alphabet, heuristic } = require ("../../src")

let permutations = alphabet ( 7 )
let result = heuristic ( permutations )
console.log ( result.length, result )
