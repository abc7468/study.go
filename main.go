package main

type binOp func(int, int) int

var a binOp

func main() {
	a = func(i1, i2 int) int {

	}
}
