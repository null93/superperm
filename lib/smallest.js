"use strict"

function smallest ( array ) {
	return array.reduce ( ( l, r ) => l.length <= r.length ? l : r )
}

module.exports = smallest
