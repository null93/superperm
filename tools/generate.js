#!/usr/bin/env node

const fs = require ("fs")
const path = require ("path")
const { alphabet, heuristic } = require ("../src")

const DIST_DIR = path.resolve ( path.dirname ( __dirname ), "dist" )

console.log (
	"| **\\|alphabet\\|** | **\\|superpermutation\\|** | **timing (ms)** |",
	"\n|:----------------:|:------------------------:|:---------------:|"
)

for ( let i = 1; i <= 9; i++ ) {
	let start = process.hrtime ()
	let permutations = alphabet ( i )
	let result = heuristic ( permutations )
	let stop = process.hrtime ( start )
	fs.writeFileSync ( path.join (`${DIST_DIR}/${i}`), result )
	let n = `${i}`
	let length = `${result.length}`
	console.log (
		"| %s | %s | %s |",
		n.padEnd ( 8 + ( n.length / 2 ), " " ).padStart ( 16, " " ),
		length.padEnd ( 12 + ( length.length / 2 ), " " ).padStart ( 24, " " ),
		`${stop [ 0 ] * 1000 + stop [ 1 ] / 1000000}`.padStart ( 15, " " )
	)
}
