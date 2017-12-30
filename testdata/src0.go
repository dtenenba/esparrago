package main

import "C"

// Without third-party trickery we can't disable the golint warning
// produced by the comment below.

//export DoubleIt
func DoubleIt(x int) int {
	return x * 2
}

func main() {}
