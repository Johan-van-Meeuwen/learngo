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
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	var (
		count = len(os.Args)
		a     = os.Args[1]
		b     = os.Args[2]
		c     = os.Args[3]
		d     = os.Args[4]
		e     = os.Args[5]
	)

	fmt.Println("There are", count, "people")
	fmt.Println("Hello great", a, "!")
	fmt.Println("Hello great", b, "!")
	fmt.Println("Hello great", c, "!")
	fmt.Println("Hello great", d, "!")
	fmt.Println("Hello great", e, "!")
	fmt.Println("Nice to meet you all.")
}
