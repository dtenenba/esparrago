package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

  SEXP IntegerVectorFromGoSlice( void* data, int ) ;
*/
import "C"

import "unsafe"

//export SumInt
func SumInt(x []int32) int32 {
	var sum int32
	for _, v := range x {
		sum += v
	}
	return sum
}

//export SumDouble
func SumDouble(x []float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum
}

//export Numbers
func Numbers(n int32) C.SEXP {
	// call a go function and get a slice

	res := make([]int32, n)
	for i := int32(0); i < n; i++ {
		res[i] = 2 * i
	}

	// handle the raw data from the slice to the C side and let it build an
	// R object from it
	return C.IntegerVectorFromGoSlice(unsafe.Pointer(&res[0]), C.int(len(res)))
}

//export DoubleIt
func DoubleIt(x int) int {
	return 2 * x
}

//export Foobar
func Foobar() string {
	return "foo" + "bar"
}

//export Nbytes
func Nbytes(s string) int {
	return len(s)
}

func main() {}
