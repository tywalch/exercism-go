package diffsquares

// SquareOfSum sums all natural numbers between one and a given int, then returns the square of that number.
func SquareOfSum(num int) int {
	// more info: https://www.mathsisfun.com/algebra/sequences-sums-arithmetic.html
	sum := (num * (1 + num)) / 2
	return sum * sum  
}

// SumOfSquares returns of a running sum of all naturual numbers between one and a given int after they have been squared. 
func SumOfSquares(num int) int {
	// more info: https://youtu.be/i7iKLZQ-vCk
	return num * (num + 1) * (2 * num + 1) / 6
}

// Finds the difference between the Square of Sums and Sum of Squares for a given int.
func Difference(num int) int {
	// return num * (num + 1) * (3 * num * num - num - 2) / 12
	return SquareOfSum(num) - SumOfSquares(num)
}

