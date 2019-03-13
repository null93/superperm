"use strict"

function flatten ( array ) {
	let result = array.shift ()
	let length = array.length
	for ( let i = 0; i < length; i++ ) {
		let current = array [ i ]
		if ( result.indexOf ( current ) > -1 ) {
			result += current
			continue
		}
		let lhs = current.slice ( 0, -1 )
		let rhs = current.slice ( -1 )
		while ( lhs.length > 0 ) {
			if ( result.endsWith ( lhs ) ) {
				break
			}
			rhs = lhs.slice ( -1 ) + rhs
			lhs = lhs.slice ( 0, -1 )
		}
		result += rhs
	}
	return result
}

module.exports = flatten
