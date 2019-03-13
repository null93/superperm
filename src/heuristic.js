"use strict"

const flatten = require ("../lib/flatten")

module.exports = permutations => {
	let current = permutations.shift ()
	let ordered = [ current ]
	while ( permutations.length > 0 ) {
		let lhs = current.slice ( 0, 1 )
		let rhs = current.slice ( 1 )
		while ( rhs.length > 0 ) {
			let searchIndex = permutations.findIndex ( e => e.startsWith ( rhs ) )
			if ( searchIndex > -1 ) {
				let found = permutations.splice ( searchIndex, 1 ) [ 0 ]
				ordered.push ( found )
				current = found
				break
			}
			lhs += rhs.slice ( 0, 1 )
			rhs = rhs.slice ( 1 )
		}
		// if ( permutations.length % 1000 === 0 ) {
		// 	console.log ( permutations.length )
		// }
	}
	return flatten ( ordered )
}
