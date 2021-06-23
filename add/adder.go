//Package adder provides a function for adding numbers
package adder

// Add takes integers and returns the sum of them.
func Add(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
