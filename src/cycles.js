"use strict"

module.exports = permutations => {
	let search = value => permutations.findIndex ( e => e.startsWith ( value ) )
	let current = permutations.shift ()
	let cycles = []
	while ( permutations.length > 0 ) {
		let cycle = [ current ]
		let foundIndex = 0
		while ( foundIndex > -1 ) {
			foundIndex = search ( current.slice ( 1 ) )
			if ( foundIndex < 0 ) continue
			current = permutations.splice ( foundIndex, 1 ) [ 0 ]
			cycle.push ( current )
		}
		cycles.push ( cycle )
		current = permutations.shift ()
	}
	return cycles
}
