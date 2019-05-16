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

func main() {
	var rows int
	fmt.Print("How many lines of Pascal's Triange would you like printed?")
	fmt.Scan(&rows)

	if rows == 0 {
		fmt.Println("0 rows is no triange")
	}
	for j := 0; j < rows; j++ { //this for cycles through rows.  J would be the row counter
		for s := rows - j; s > 0; s-- { //this is to enter correct spacing, to make it look like a triange
			fmt.Print(" ")
		}

		for m := 0; m <= j; m++ { //m should be the position in the row
			fmt.Printf("%v ", nChooseK(j, m))

		}
		fmt.Println()
	}
}
