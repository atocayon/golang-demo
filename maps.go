package main

import (
	"fmt"
	"maps"
)

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
func mapsdemo() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	// Output:
	// map: map[k1:7 k2:13]
	// v1: 7
	// v3: 0
	// len: 2
	// map: map[k1:7]
	// map: map[]
	// prs: false
	// map: map[bar:2 foo:1]
	// n == n2
}
