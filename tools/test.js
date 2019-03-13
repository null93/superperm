#!/usr/bin/env node

const fs = require ("fs")
const path = require ("path")
const { alphabet, heuristic } = require ("../src")

const PASS_TPL = "\033[33m%s: \033[32mpassed\033[0m"
const FAIL_TPL = "\033[33m%s: \033[31mfailed\033[0m (%d missing)"
const DIST_DIR = path.resolve ( path.dirname ( __dirname ), "dist" )
const files = fs.readdirSync ( DIST_DIR ).map ( e => path.join ( DIST_DIR, e ) )

for ( let file of files ) {
	let cardinality = parseInt ( path.basename ( file ) )
	let permutations = alphabet ( cardinality )
	let result = fs.readFileSync ( file ).toString ()
	let excluded = permutations.filter ( e => result.indexOf ( e ) < 0 )
	if ( excluded.length > 0 ) {
		console.error ( FAIL_TPL, cardinality, excluded.length )
	}
	else {
		console.log ( PASS_TPL, cardinality )
	}
}
