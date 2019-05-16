package main

import "fmt"

func factorial(n int) uint64 {
	var factVal uint64 = 1
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	default:
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
		return factVal
	}
}

// n choose k
//(n,k) = n!/(k!(n-k)!)
// n = row
// k = position

func nChooseK(n, k int) uint64 {
	var x uint64
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
	for j := 0; j <= rows; j++ { //this for cycles through rows.  J would be the row counter
		for s := rows - j; s > 0; s-- { //this is to enter correct spacing, to make it look like a triange
			print(" ")
		}

		for m := 0; m <= j; m++ { //m should be the position in the row
			print(nChooseK(j, m))

		}
		print("\n")
	}
}
