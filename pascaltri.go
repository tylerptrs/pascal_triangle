package main

import "fmt"

var factVal uint64 = 1
var i, n, k int
var x uint64

func factorial(n int) uint64 {
	for i := 1; i <= n; i++ {
		factVal *= uint64(i)
	}
	return factVal
}

func nChooseK(n, k int) uint64 {
	x = factorial(n) / (factorial(k) * factorial(n-k))
	return x
}

var rows int

func main() {
	print("How many lines of Pascal's Triange would you like printed?")
	fmt.Scan(&rows)
	rows = rows - 1
	if rows == 0 {
		print("0 rows is no triange")
	}
	for j := 0; j <= rows; j++ {
		for s := rows - j; s > 0; s-- {
			print(" ")
		}
		print(nChooseK(row))
	}
}

// n choose k
//(n,k) = n!/(k!(n-k)!)
// n = row
// k = position
