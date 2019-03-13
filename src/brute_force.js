"use strict"

const permute = require ("heaps-permute")
const flatten = require ("../lib/flatten")
const smallest = require ("../lib/smallest")

module.exports = permutations => {
	let orders = permute ( permutations ).map ( e => flatten ( e ) )
	return smallest ( orders )
}
