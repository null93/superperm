"use strict"

const permute = require ("heaps-permute")

function alphabet ( size, alpha = false ) {
	let set = new Array ( size ).fill ("")
	let alphaInc = i => String.fromCharCode ( "A".charCodeAt ( 0 ) + i )
	let numInc = i => i
	let increment = alpha ? alphaInc : numInc
	let variations = set.map ( ( e, i ) => `${increment ( i + 1 )}` )
	let permutations = permute ( variations )
	return permutations.map ( e => e.join ("") )
}

module.exports = alphabet
