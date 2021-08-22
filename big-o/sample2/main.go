package sample2

func anotherFunc() {

}

// Big - O Calc
// BIG O -> 3 + 4n(n + n + n)
func funCallenge(input []int) int {
	a := 10    // O(1)
	a = 50 + 3 // O(1)

	for _, _ = range input { // O(n)
		anotherFunc() // O(n)
		_ = true      // O(n)
		a++           // O(n)
	}
	return a // O(1)
}
