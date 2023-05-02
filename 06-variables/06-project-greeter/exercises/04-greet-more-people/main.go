// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	if len(os.Args) >= 4 {
		var (
			count = len(os.Args) - 1
			bilbo = os.Args[1]
			balbo = os.Args[2]
			bungo = os.Args[3]
		)

		fmt.Println("There are", count, "people")
		fmt.Println("Hello great", bilbo)
		fmt.Println("Hello great", balbo)
		fmt.Println("Hello great", bungo)
		fmt.Println("Nice to meet you all.")
	}

	if len(os.Args) <= 3 {
		fmt.Println("Please give more names!")
	}
}
